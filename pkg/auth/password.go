package auth

import (
	"fmt"
	"strings"
	"unicode"
)

// InvalidPasswordError is raised when a password does not pass validation.
type InvalidPasswordError struct {
	Issues []string `json:"issues"`
}

func (e InvalidPasswordError) Error() string {
	return strings.Join(e.Issues, ", ")
}

// IPasswordService defines all functions required for managing passwords.
type IPasswordService interface {
	ValidatePassword(password string) error
	VerifyPassword(password string, passwordHash string) (bool, error)

	GeneratePasswordHash(password string) (string, error)
}

// PasswordConfig defines all fields required to create a PasswordService.
type PasswordConfig struct {
	EncryptRepo IEncryptRepo

	MinLength        int
	MaxLength        int
	RequireUppercase bool
	RequireLowercase bool
	RequireNumber    bool
	RequireSpecial   bool
}

// PasswordService implements a standard password managing service.
type PasswordService struct {
	encryptRepo IEncryptRepo

	minLength        int
	maxLength        int
	requireUppercase bool
	requireLowercase bool
	requireNumber    bool
	requireSpecial   bool
}

// NewPasswordService creates a new PasswordService.
func NewPasswordService(config PasswordConfig) *PasswordService {
	return &PasswordService{
		encryptRepo:      config.EncryptRepo,
		minLength:        config.MinLength,
		maxLength:        config.MaxLength,
		requireUppercase: config.RequireUppercase,
		requireLowercase: config.RequireLowercase,
		requireNumber:    config.RequireNumber,
		requireSpecial:   config.RequireSpecial,
	}
}

// ValidatePassword checks a given password against any enabled complexity rules.
func (s *PasswordService) ValidatePassword(password string) error {
	hasNumber := false
	hasUppercase := false
	hasLowercase := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		default:
		}
	}

	issues := []string{}
	if len(password) < s.minLength {
		issues = append(issues, fmt.Sprintf("password must be at least %d characters", s.minLength))
	}

	if s.requireUppercase && !hasUppercase {
		issues = append(issues, "at least one uppercase character required")
	}

	if s.requireLowercase && !hasLowercase {
		issues = append(issues, "at least one lowercase character required")
	}

	if s.requireNumber && !hasNumber {
		issues = append(issues, "at least one numeric character required")
	}

	if s.requireSpecial && !hasSpecial {
		issues = append(issues, "at least one special character required")
	}

	if len(issues) > 0 {
		return InvalidPasswordError{
			Issues: issues,
		}
	}

	return nil
}

// VerifyPassword compares a password to a hashed password.
func (s *PasswordService) VerifyPassword(password string, passwordHash string) (bool, error) {
	result, err := s.encryptRepo.Verify(password, passwordHash)
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}

	return result, nil
}

// GeneratePasswordHash encrypts the given password.
func (s *PasswordService) GeneratePasswordHash(password string) (string, error) {
	result, err := s.encryptRepo.Generate(password)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return result, nil
}
