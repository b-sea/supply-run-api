package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/b-sea/go-auth/password"
	"github.com/b-sea/go-auth/token"
	"github.com/sirupsen/logrus"
)

type (
	contextKey int
	tokenType  string
)

const (
	tokenAccessAud  tokenType = "access"
	tokenRefreshAud tokenType = "refresh"
	headerKey                 = "Authorization"
	headerTokenType           = "Bearer "
	contextUserKey            = contextKey(1)
)

type Repository interface {
	GetUser(username string) (*User, error)
}

type Service struct {
	repo     Repository
	token    *token.Service
	password *password.Service
}

func NewService(r Repository, t *token.Service, p *password.Service) *Service {
	return &Service{
		repo:     r,
		token:    t,
		password: p,
	}
}

func (s *Service) Token() *token.Service {
	return s.token
}

func (s *Service) Password() *password.Service {
	return s.password
}

func (s *Service) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header[headerKey]
		if authHeader == nil || len(authHeader) != 1 {
			logrus.Errorf(
				"%s: invalid or missing %s header found",
				http.StatusText(http.StatusUnauthorized),
				headerKey,
			)
			jsonError(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		tokenString, ok := strings.CutPrefix(authHeader[0], headerTokenType)
		if !ok {
			logrus.Errorf("%s: incorrect token format", http.StatusText(http.StatusUnauthorized))
			jsonError(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		jwt, err := s.token.ParseAccessToken(tokenString)
		if err != nil {
			logrus.Errorf("%s: %s", http.StatusText(http.StatusUnauthorized), err)
			jsonError(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		sub, err := jwt.Claims.GetSubject()
		if err != nil {
			logrus.Errorf("%s: %s", http.StatusText(http.StatusUnauthorized), err)
			jsonError(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		user, err := s.repo.GetUser(sub)
		if err != nil {
			logrus.Errorf("%s: %s", http.StatusText(http.StatusUnauthorized), err)
			jsonError(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), contextUserKey, user)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

func FromContext(ctx context.Context) *User {
	user, _ := ctx.Value(contextUserKey).(*User)
	return user
}

func jsonError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
