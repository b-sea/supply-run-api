package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/b-sea/supply-run-api/internal/repository"
	"github.com/b-sea/supply-run-api/pkg/auth"
)

const (
	basicAuthSep = ":"
	basicAuthLen = 2
)

type IAccountService interface {
	BasicAuth(encoded string) (*model.TokenSet, error)
	RefreshToken(ctx context.Context, token string) (string, error)
	ChangePassword(ctx context.Context, password string) (bool, error)

	GetOne(ctx context.Context, id model.ID) (*model.Account, error)
	Create(ctx context.Context, input model.CreateAccountInput) (model.CreateResult, error)
	Update(ctx context.Context, input model.UpdateAccountInput) (model.UpdateResult, error)
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

func (s *AccountService) BasicAuth(encoded string) (*model.TokenSet, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, model.BadRequestError{Reason: "invalid credential format"}
	}

	credentials := strings.Split(string(decoded), basicAuthSep)
	if len(credentials) != basicAuthLen {
		return nil, model.BadRequestError{Reason: "invalid credential format"}
	}

	found, err := s.repo.Find(&model.AccountFilter{Username: model.StringFilter{Eq: &credentials[0]}})
	if err != nil || len(found) != 1 {
		return nil, model.AuthenticationError{}
	}

	if ok, err := s.pwd.VerifyPassword(found[0].Password, credentials[1]); !ok || err != nil {
		return nil, model.AuthenticationError{}
	}

	found[0].Password, err = s.pwd.GeneratePasswordHash(credentials[1])
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	access, err := s.token.GenerateAccessToken(found[0].ID.String())
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	refresh, err := s.token.GenerateRefreshToken(found[0].ID.String())
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &model.TokenSet{
		Access:  access,
		Refresh: refresh,
	}, nil
}
