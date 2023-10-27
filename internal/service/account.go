// Package service implements all business logic for the API.
package service

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/b-sea/supply-run-api/internal/repository"
	"github.com/b-sea/supply-run-api/pkg/auth"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type contextKey string

const accountIDContextKey contextKey = "accountID"

type IAccount interface {
	Signup(input model.SignupInput) error
	Login(email string, password string) (*model.LoginResult, error)
	RefreshToken(token string) (string, error)

	Profile(ctx context.Context) (*model.Account, error)
	Update(ctx context.Context, input model.UpdateAccountInput) error
	Delete(ctx context.Context, id model.ID) error

	NewContext(ctx context.Context, email string) (context.Context, error)
}

type AccountConfig struct {
	Password auth.IPasswordService
	Token    auth.ITokenService
	Repo     repository.IAccount
}

type Account struct {
	pwd   auth.IPasswordService
	token auth.ITokenService
	repo  repository.IAccount
}

func NewAccount(config AccountConfig) *Account {
	return &Account{
		pwd:   config.Password,
		token: config.Token,
		repo:  config.Repo,
	}
}

func (s Account) Signup(input model.SignupInput) error {
	if _, err := mail.ParseAddress(input.Email); err != nil {
		return model.ValidationError{
			Issues: []string{"invalid email format"},
		}
	}

	found, err := s.repo.GetByEmail(input.Email)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("%w", err)
	}
	if found != nil {
		return model.ValidationError{
			Issues: []string{"an account with that email already exists"},
		}
	}

	if err := s.pwd.ValidatePassword(input.Password); err != nil {
		var target auth.InvalidPasswordError
		if errors.As(err, &target) {
			return model.ValidationError{
				Issues: target.Issues,
			}
		}

		logrus.Error(err)
		return fmt.Errorf("%w", err)
	}

	account := input.ToNode(uuid.NewString(), time.Now().UTC())
	account.Password, err = s.pwd.GeneratePasswordHash(input.Password)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("%w", err)
	}

	if err := s.repo.Create(account); err != nil {
		logrus.Error(err)
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (s *Account) Login(email string, password string) (*model.LoginResult, error) {
	found, err := s.repo.GetByEmail(email)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	if found != nil {
		return nil, ErrAuthentication
	}

	if ok, err := s.pwd.VerifyPassword(password, found.Password); !ok || err != nil {
		return nil, ErrAuthentication
	}

	now := time.Now().UTC()
	found.LastLogin = &now
	found.Password, err = s.pwd.GeneratePasswordHash(password)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if err = s.repo.Update(*found); err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	access, err := s.token.GenerateAccessToken(found.Email)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	refresh, err := s.token.GenerateRefreshToken(found.Email)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	return &model.LoginResult{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *Account) RefreshToken(token string) (string, error) {
	refreshToken, err := s.token.ParseRefreshToken(token)
	if err != nil {
		logrus.Error(err)
		return "", ErrAuthentication
	}

	subject, _ := refreshToken.Claims.GetSubject()
	access, err := s.token.GenerateAccessToken(subject)
	if err != nil {
		logrus.Error(err)
		return "", fmt.Errorf("%w", err)
	}

	return access, nil
}

func (s *Account) Profile(ctx context.Context) (*model.Account, error) {
	accountID, ok := AccountIDFromContext(ctx)
	if !ok {
		return nil, ErrAuthentication
	}

	account, err := s.repo.GetByID(accountID.Key)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	if account == nil {
		return nil, model.NotFoundError{ID: *accountID}
	}

	return account, nil
}

func (s *Account) Update(ctx context.Context, input model.UpdateAccountInput) error {
	accountID, ok := AccountIDFromContext(ctx)
	if !ok {
		return ErrAuthentication
	}

	account, err := s.repo.GetByID(accountID.Key)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("%w", err)
	}

	input.MergeNode(account, time.Now().UTC())

	if input.Password != nil {
		if err := s.pwd.ValidatePassword(*input.Password); err != nil {
			var target auth.InvalidPasswordError
			if errors.As(err, &target) {
				return model.ValidationError{
					Issues: target.Issues,
				}
			}

			logrus.Error(err)
			return fmt.Errorf("%w", err)
		}

		if account.Password, err = s.pwd.GeneratePasswordHash(*input.Password); err != nil {
			logrus.Error(err)
			return fmt.Errorf("%w", err)
		}
	}

	if err := s.repo.Update(*account); err != nil {
		logrus.Error(err)
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (s *Account) Delete(ctx context.Context, id model.ID) error {
	accountID, ok := AccountIDFromContext(ctx)
	if !ok {
		return ErrAuthentication
	}

	if err := s.repo.Delete(accountID.Key); err != nil {
		logrus.Error(err)
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (s *Account) NewContext(ctx context.Context, email string) (context.Context, error) {
	found, err := s.repo.GetByEmail(email)
	if err != nil {
		logrus.Error(err)
		return ctx, fmt.Errorf("%w", err)
	}
	if found == nil {
		return ctx, ErrAuthentication
	}
	if !found.IsVerified {
		return ctx, model.ValidationError{Issues: []string{"account not verified"}}
	}

	return context.WithValue(ctx, accountIDContextKey, &found.ID), nil
}

func AccountIDFromContext(ctx context.Context) (*model.ID, bool) {
	id, ok := ctx.Value(accountIDContextKey).(*model.ID)
	if !ok || id == nil {
		return nil, false
	}

	return id, true
}
