// Package auth is responsible for managing authentication and authorization.
package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type contextKey int

const (
	tokenAccessAud  = "access"
	tokenRefreshAud = "refresh"

	headerKey       = "Authorization"
	headerTokenType = "Bearer "
	contextTokenKey = contextKey(1)
)

var (
	// Timestamp is a function to generate a timestamp value.
	Timestamp = time.Now //nolint: gochecknoglobals

	// ErrRSAKey is raised when RSA keys encounter an error.
	ErrRSAKey = errors.New("rsa key error")

	errJWTSigning = errors.New("jwt signing error")
)

func rsaKeyError(value interface{}) error {
	return fmt.Errorf("%w: %v", ErrRSAKey, value)
}

func jwtSigningError(value interface{}) error {
	return fmt.Errorf("%w: %v", errJWTSigning, value)
}

// ITokenService defines all functions required for managing auth tokens.
type ITokenService interface {
	ParseAccessToken(tokenString string) (*jwt.Token, error)
	ParseRefreshToken(tokenString string) (*jwt.Token, error)

	GenerateAccessToken(sub string) (string, error)
	GenerateRefreshToken(sub string) (string, error)

	FromHeader(header http.Header) (string, bool)
}

// TokenConfig defines all fields required to create a TokenService.
type TokenConfig struct {
	SignMethod string
	PublicKey  []byte
	PrivateKey []byte

	Issuer         string
	Audience       string
	AccessTimeout  time.Duration
	RefreshTimeout time.Duration

	IDGenerator func() string
}

// TokenService implements a standard JWT auth service.
type TokenService struct {
	signMethod string
	signKey    *rsa.PrivateKey
	verifyKey  *rsa.PublicKey

	issuer         string
	audience       string
	accessTimeout  time.Duration
	refreshTimeout time.Duration

	idGenerator func() string
}

// NewTokenService creates a new TokenService.
func NewTokenService(config TokenConfig) (*TokenService, error) {
	if config.SignMethod != "RS256" {
		return nil, rsaKeyError("only sign method RS256 is currently supported")
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(config.PublicKey)
	if err != nil {
		return nil, rsaKeyError(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(config.PrivateKey)
	if err != nil {
		return nil, rsaKeyError(err)
	}

	if config.IDGenerator == nil {
		config.IDGenerator = uuid.NewString
	}

	return &TokenService{
		signMethod:     config.SignMethod,
		verifyKey:      publicKey,
		signKey:        privateKey,
		issuer:         config.Issuer,
		audience:       config.Audience,
		accessTimeout:  config.AccessTimeout,
		refreshTimeout: config.RefreshTimeout,
		idGenerator:    config.IDGenerator,
	}, nil
}

// ParseAccessToken verifies and transforms a given token string into an access JWT.
func (s *TokenService) ParseAccessToken(tokenString string) (*jwt.Token, error) {
	var claims jwt.RegisteredClaims

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			return s.verifyKey, nil
		},
		jwt.WithIssuer(s.issuer),
		jwt.WithAudience(s.audience),
		jwt.WithAudience(tokenAccessAud),
		jwt.WithIssuedAt(),
		jwt.WithValidMethods([]string{s.signMethod}),
	)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return token, nil
}

// ParseRefreshToken verifies and transforms a given token string into a refresh JWT.
func (s *TokenService) ParseRefreshToken(tokenString string) (*jwt.Token, error) {
	var claims jwt.RegisteredClaims

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			return s.verifyKey, nil
		},
		jwt.WithIssuer(s.issuer),
		jwt.WithAudience(s.audience),
		jwt.WithAudience(tokenRefreshAud),
		jwt.WithIssuedAt(),
		jwt.WithValidMethods([]string{s.signMethod}),
	)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return token, nil
}

// GenerateAccessToken creates and signes a new access JWT.
func (s *TokenService) GenerateAccessToken(sub string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.GetSigningMethod(s.signMethod),
		&jwt.RegisteredClaims{
			ID:        s.idGenerator(),
			Subject:   sub,
			Issuer:    s.issuer,
			Audience:  jwt.ClaimStrings([]string{s.audience, tokenAccessAud}),
			ExpiresAt: jwt.NewNumericDate(Timestamp().Add(s.accessTimeout)),
			IssuedAt:  jwt.NewNumericDate(Timestamp()),
		},
	)

	signed, err := token.SignedString(s.signKey)
	if err != nil {
		return "", jwtSigningError(err)
	}

	return signed, nil
}

// GenerateRefreshToken creates and signes a new refresh JWT.
func (s *TokenService) GenerateRefreshToken(sub string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.GetSigningMethod(s.signMethod),
		&jwt.RegisteredClaims{
			ID:        s.idGenerator(),
			Subject:   sub,
			Issuer:    s.issuer,
			Audience:  jwt.ClaimStrings([]string{s.audience, tokenRefreshAud}),
			ExpiresAt: jwt.NewNumericDate(Timestamp().Add(s.accessTimeout)),
			IssuedAt:  jwt.NewNumericDate(Timestamp()),
		},
	)

	signed, err := token.SignedString(s.signKey)
	if err != nil {
		return "", jwtSigningError(err)
	}

	return signed, nil
}

// FromHeader retrieves a token string from the given headers, if it exists.
func (s *TokenService) FromHeader(header http.Header) (string, bool) {
	bearer := header[headerKey]
	if bearer == nil || len(bearer) != 1 {
		return "", false
	}

	token, ok := strings.CutPrefix(bearer[0], headerTokenType)
	if !ok {
		return "", false
	}

	return token, true
}
