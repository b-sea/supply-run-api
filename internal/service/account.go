// Package service implements all business logic for the API.
package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/b-sea/supply-run-api/internal/repository"
	"github.com/b-sea/supply-run-api/pkg/auth"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type contextKey string

const accountIDContextKey contextKey = "accountID"

type IAccountService interface {
	Signup(input model.CreateAccountInput) (*model.ID, error)
	Login(email string, password string) (*model.LoginResult, error)
	RefreshToken(token string) (string, error)
	NewContext(ctx context.Context, email string) context.Context
	Profile(ctx context.Context) (*model.Account, error)
	Update(ctx context.Context, input model.UpdateAccountInput) error
	Delete(ctx context.Context, id model.ID) error
}

type AccountConfig struct {
	Password auth.IPasswordService
	Token    auth.ITokenService
	Repo     repository.IAccountRepo
}

type AccountService struct {
	pwd   auth.IPasswordService
	token auth.ITokenService
	repo  repository.IAccountRepo
}

func NewAccountService(config AccountConfig) *AccountService {
	return &AccountService{
		pwd:   config.Password,
		token: config.Token,
		repo:  config.Repo,
	}
}

func (s *AccountService) Login(email string, password string) (*model.LoginResult, error) {
	found, err := s.repo.Find(&model.AccountFilter{Email: &model.StringFilter{Eq: &email}})
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	if len(found) != 1 {
		return nil, ErrAuthentication
	}

	if ok, err := s.pwd.VerifyPassword(password, found[0].Password); !ok || err != nil {
		return nil, ErrAuthentication
	}

	now := time.Now().UTC()
	found[0].LastLogin = &now
	found[0].Password, err = s.pwd.GeneratePasswordHash(password)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if err = s.repo.Update(*found[0]); err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	access, err := s.token.GenerateAccessToken(found[0].Email)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	refresh, err := s.token.GenerateRefreshToken(found[0].Email)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	return &model.LoginResult{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *AccountService) RefreshToken(token string) (string, error) {
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

func (s AccountService) Signup(input model.CreateAccountInput) (*model.ID, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	found, err := s.repo.Find(&model.AccountFilter{Email: &model.StringFilter{Eq: &input.Email}})
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}
	if len(found) != 0 {
		return nil, model.ValidationError{
			Issues: []string{"an account with that email already exists"},
		}
	}

	if err := s.pwd.ValidatePassword(input.Password); err != nil {
		var target auth.InvalidPasswordError
		if errors.As(err, &target) {
			return nil, model.ValidationError{
				Issues: target.Issues,
			}
		}

		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	entity := input.ToEntity(uuid.NewString(), time.Now().UTC())
	entity.Password, err = s.pwd.GeneratePasswordHash(input.Password)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	if err := s.repo.Create(entity); err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	return &entity.ID, nil
}

func (s *AccountService) Profile(ctx context.Context) (*model.Account, error) {
	accountID, ok := AccountIDFromContext(ctx)
	if !ok {
		return nil, ErrAuthentication
	}

	account, err := s.repo.GetOne(accountID.Key)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	return account, nil
}

func (s *AccountService) Update(ctx context.Context, input model.UpdateAccountInput) error {
	accountID, ok := AccountIDFromContext(ctx)
	if !ok {
		return ErrAuthentication
	}

	account, err := s.repo.GetOne(accountID.Key)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("%w", err)
	}

	input.MergeEntity(account, time.Now().UTC())

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

func (s *AccountService) Delete(ctx context.Context, id model.ID) error {
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

func (s *AccountService) NewContext(ctx context.Context, email string) context.Context {
	found, err := s.repo.Find(&model.AccountFilter{Email: &model.StringFilter{Eq: &email}})
	if err != nil {
		logrus.Error(err)
		return ctx
	}
	if len(found) != 1 {
		return ctx
	}

	return context.WithValue(ctx, accountIDContextKey, &found[0].ID)
}

func AccountIDFromContext(ctx context.Context) (*model.ID, bool) {
	id, ok := ctx.Value(accountIDContextKey).(*model.ID)
	if !ok || id == nil {
		return nil, false
	}

	return id, true
}
