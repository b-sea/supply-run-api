// Package auth defines authentication workflows and entities.
package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/b-sea/go-auth/password"
	"github.com/b-sea/go-auth/token"
	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/gorilla/mux"
)

const (
	authHeaderKey = "Authorization"
	bearerPrefix  = "Bearer "

	userKey contextKey = "user"
)

type contextKey string

// Recorder defines functions for tracking auth-based metrics.
type Recorder interface {
	RequestAuthorized(username string)
}

// Service is an authentication service.
type Service struct {
	token     *token.Service
	pwd       *password.Service
	repo      Repository
	recorder  Recorder
	timestamp func() time.Time
}

// NewService creates a new authentication Service.
func NewService(repo Repository, token *token.Service, pwd *password.Service, recorder Recorder) *Service {
	return &Service{
		token:     token,
		pwd:       pwd,
		repo:      repo,
		recorder:  recorder,
		timestamp: time.Now,
	}
}

// TokensCommand is the input for Tokens.
type TokensCommand struct {
	Username string
	Password string
}

// Tokens generates access and refresh tokens from credentials.
func (s *Service) Tokens(_ context.Context, cmd TokensCommand) (*Tokens, error) {
	user, err := s.repo.GetAuthUser(cmd.Username)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrUnauthorized
		}

		return nil, err
	}

	account, err := s.repo.GetAccount(user.id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrUnauthorized
		}

		return nil, err
	}

	if ok, err := s.pwd.Verify(cmd.Password, account.password); err != nil || !ok {
		return nil, ErrUnauthorized
	}

	hashed, err := s.pwd.GenerateHash(cmd.Password)
	if err != nil {
		return nil, ErrUnauthorized
	}

	if err := account.Update(s.timestamp().UTC(), setPassword(hashed)); err != nil {
		return nil, err
	}

	if err := s.repo.UpdateAccount(account); err != nil {
		return nil, err
	}

	access, err := s.token.GenerateAccessToken(account.username)
	if err != nil {
		return nil, err
	}

	refresh, err := s.token.GenerateRefreshToken(account.username)
	if err != nil {
		return nil, err
	}

	s.recorder.RequestAuthorized(user.username)

	return &Tokens{
		Access:  access,
		Refresh: refresh,
	}, nil
}

// GetAccount returns the user Account of an authenticated user.
func (s *Service) GetAccount(ctx context.Context) (*Account, error) {
	user, err := FromContext(ctx)
	if err != nil {
		return nil, err
	}

	account, err := s.repo.GetAccount(user.id)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// CreateAccountCommand is the input for CreateAccount.
type CreateAccountCommand struct {
	Username string
	Password string
}

// CreateAccount creates a new user Account.
func (s *Service) CreateAccount(_ context.Context, cmd CreateAccountCommand) error {
	// TODO: Validate username length
	if _, err := s.repo.GetAuthUser(cmd.Username); err != nil {
		if !errors.Is(err, ErrNotFound) {
			return err
		}
	} else {
		return ErrDuplicateAccount
	}

	if err := s.pwd.ValidatePassword(cmd.Password); err != nil {
		return err
	}

	hashed, err := s.pwd.GenerateHash(cmd.Password)
	if err != nil {
		return err
	}

	account := NewAccount(entity.NewID(), cmd.Username, hashed, s.timestamp().UTC())

	if err := s.repo.CreateAccount(account); err != nil {
		return err
	}

	return nil
}

// UpdateAccountCommand is the input for UpdateAccount.
type UpdateAccountCommand struct {
	Password *string
}

// UpdateAccount updates an existing user Account.
func (s *Service) UpdateAccount(ctx context.Context, cmd UpdateAccountCommand) error {
	user, err := FromContext(ctx)
	if err != nil {
		return err
	}

	options := make([]AccountOption, 0)

	if cmd.Password != nil {
		if err := s.pwd.ValidatePassword(*cmd.Password); err != nil {
			return err
		}

		hashed, err := s.pwd.GenerateHash(*cmd.Password)
		if err != nil {
			return err
		}

		options = append(options, setPassword(hashed))
	}

	if len(options) == 0 {
		return nil
	}

	account, err := s.repo.GetAccount(user.id)
	if err != nil {
		return err
	}

	if err := account.Update(s.timestamp().UTC(), options...); err != nil {
		return err
	}

	if err := s.repo.UpdateAccount(account); err != nil {
		return err
	}

	return nil
}

// DeleteAccount deletes an existing user Account.
func (s *Service) DeleteAccount(ctx context.Context) error {
	user, err := FromContext(ctx)
	if err != nil {
		return err
	}

	if err := s.repo.DeleteAccount(user.id); err != nil {
		return err
	}

	return nil
}

// Middleware is authentication middleware.
//
// This will extract token data from a request "Authentication" header and store the
// associated authenticated user in the request context.
// If no token is found, this will not error. The request context will simply not
// contain an authenticated user.
func (s *Service) Middleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			header := request.Header.Get(authHeaderKey)

			if !strings.HasPrefix(header, bearerPrefix) {
				next.ServeHTTP(writer, request)
				return
			}

			jwt, err := s.token.ParseAccessToken(strings.TrimPrefix(header, bearerPrefix))
			if err != nil {
				next.ServeHTTP(writer, request)
				return
			}

			subject, err := jwt.Claims.GetSubject()
			if err != nil {
				next.ServeHTTP(writer, request)
				return
			}

			user, err := s.repo.GetAuthUser(subject)
			if err != nil {
				next.ServeHTTP(writer, request)
				return
			}

			s.recorder.RequestAuthorized(user.username)
			next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), userKey, user)))
		})
	}
}

// FromContext retrieves an authenticated User from a context.
func FromContext(ctx context.Context) (*User, error) {
	user, ok := ctx.Value(userKey).(*User)
	if !ok {
		return nil, ErrUnauthorized
	}

	return user, nil
}

// Challenge configures an HTTP response with appropriate header values when tokens are incorrect or missing.
func Challenge(writer http.ResponseWriter) {
	writer.Header().Add("WWW-Authenticate", strings.Trim(bearerPrefix, " "))
	http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}
