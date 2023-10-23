// Package auth is responsible for managing authentication and authorization.
package auth

import (
	"context"
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
	tokenTypeKey     = "typ"
	tokenTypeAccess  = "JWT-ACCESS"
	tokenTypeRefresh = "JWT-REFRESH"

	headerKey       = "Authorization"
	headerTokenType = "Bearer "
	contextTokenKey = contextKey(1)
)

var (
	// Timestamp is a function to generate a timestamp value.
	Timestamp = time.Now //nolint: gochecknoglobals

	// ErrAuthentication is raised when authentication fails.
	ErrAuthentication = errors.New("authentication error")

	// ErrAuthorization is raised when authorization fails.
	ErrAuthorization = errors.New("authorization error")

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
}

// TokenConfig defines all fields required to create a TokenService.
type TokenConfig struct {
	SignMethod string
	PublicKey  []byte
	PrivateKey []byte

	Issuer         string
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
		accessTimeout:  config.AccessTimeout,
		refreshTimeout: config.RefreshTimeout,
		idGenerator:    config.IDGenerator,
	}, nil
}

// ParseAccessToken verifies and transforms a given token string into an access JWT.
func (s *TokenService) ParseAccessToken(tokenString string) (*jwt.Token, error) {
	token, err := s.parseToken(tokenString, tokenTypeAccess)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// ParseRefreshToken verifies and transforms a given token string into a refresh JWT.
func (s *TokenService) ParseRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := s.parseToken(tokenString, tokenTypeRefresh)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *TokenService) parseToken(tokenString string, typ string) (*jwt.Token, error) {
	var claims jwt.RegisteredClaims

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			return s.verifyKey, nil
		},
		jwt.WithIssuer(s.issuer),
		jwt.WithIssuedAt(),
		jwt.WithValidMethods([]string{s.signMethod}),
	)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	// Custom verification
	tokenType, ok := token.Header[tokenTypeKey].(string)
	if !ok || tokenType != typ || !token.Valid {
		return nil, fmt.Errorf("%w", jwt.ErrTokenUnverifiable)
	}

	return token, nil
}

// GenerateAccessToken creates and signes a new access JWT.
func (s *TokenService) GenerateAccessToken(sub string) (string, error) {
	token, err := s.generateToken(sub, s.accessTimeout, tokenTypeAccess)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GenerateRefreshToken creates and signes a new refresh JWT.
func (s *TokenService) GenerateRefreshToken(sub string) (string, error) {
	token, err := s.generateToken(sub, s.refreshTimeout, tokenTypeRefresh)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *TokenService) generateToken(sub string, timeout time.Duration, typ string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.GetSigningMethod(s.signMethod),
		&jwt.RegisteredClaims{
			ID:        s.idGenerator(),
			Subject:   sub,
			Issuer:    s.issuer,
			ExpiresAt: jwt.NewNumericDate(Timestamp().Add(timeout)),
			IssuedAt:  jwt.NewNumericDate(Timestamp()),
		},
	)

	// Custom headers
	token.Header[tokenTypeKey] = typ

	signed, err := token.SignedString(s.signKey)
	if err != nil {
		return "", jwtSigningError(err)
	}

	return signed, nil
}

// FromHeader retrieves a token string from the given headers, if it exists.
func FromHeader(header http.Header) (string, bool) {
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

// NewTokenContext returns a new context that carries a JWT.
func NewTokenContext(ctx context.Context, token *jwt.Token) context.Context {
	return context.WithValue(ctx, contextTokenKey, token)
}

// SubjectFromContext retrieves a token sub claim from the given context, if it exists.
func SubjectFromContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(contextTokenKey).(*jwt.Token)
	if !ok || token == nil {
		return "", false
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)

	return claims.Subject, ok
}
