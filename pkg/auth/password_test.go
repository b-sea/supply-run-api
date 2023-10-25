package auth_test

import (
	"errors"
	"testing"

	"github.com/b-sea/supply-run-api/pkg/auth"
	"github.com/stretchr/testify/assert"
)

type MockEncryptRepo struct {
	VerifyResult   bool
	VerifyErr      error
	GenerateResult string
	GenerateErr    error
}

func (r *MockEncryptRepo) Verify(_ string, _ string) (bool, error) {
	return r.VerifyResult, r.VerifyErr
}

func (r *MockEncryptRepo) Generate(_ string) (string, error) {
	return r.GenerateResult, r.GenerateErr
}

func TestPasswordServiceValidatePassword(t *testing.T) {
	t.Parallel()

	type test struct {
		password string
		err      error
	}

	testCases := map[string]test{
		"success": {
			password: "P@ssw0rd",
		},
		"missing uppercase": {
			password: "p@ssw0rd",
			err:      auth.InvalidPasswordError{},
		},
		"missing lowercase": {
			password: "P@SSW0RD",
			err:      auth.InvalidPasswordError{},
		},
		"missing number": {
			password: "P@ssword",
			err:      auth.InvalidPasswordError{},
		},
		"missing special": {
			password: "Passw0rd",
			err:      auth.InvalidPasswordError{},
		},
		"too short": {
			password: "Pwd0!",
			err:      auth.InvalidPasswordError{},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		config := auth.PasswordConfig{
			MinLength:      8,
			MaxLength:      100,
			RequireUpper:   true,
			RequireLower:   true,
			RequireNumber:  true,
			RequireSpecial: true,
		}

		pwdService := auth.NewPasswordService(config)

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			err := pwdService.ValidatePassword(testCase.password)
			if testCase.err == nil {
				assert.NoError(t, err, "no error expected")
			} else {
				assert.ErrorAs(t, err, &testCase.err, "different errors")
			}
		})
	}
}

func TestPasswordServiceVerifyPassword(t *testing.T) {
	t.Parallel()

	type test struct {
		encryptRepo auth.IEncryptRepo
		password    string
		hash        string
		result      bool
		err         error
	}

	testCases := map[string]test{
		"matched": {
			encryptRepo: &MockEncryptRepo{
				VerifyResult: true,
			},
			password: "password",
			hash:     "1a2b3c4d",
			result:   true,
		},
		"no match": {
			encryptRepo: &MockEncryptRepo{
				VerifyResult: false,
			},
			password: "password",
			hash:     "1a2b3c4d",
			result:   false,
		},
		"error": {
			encryptRepo: &MockEncryptRepo{
				VerifyErr: errors.New("some hash error"),
			},
			password: "password",
			hash:     "1a2b3c4d",
			err:      errors.New("some hash error"),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		config := auth.PasswordConfig{
			EncryptRepo:    testCase.encryptRepo,
			MinLength:      8,
			MaxLength:      100,
			RequireUpper:   true,
			RequireLower:   true,
			RequireNumber:  true,
			RequireSpecial: true,
		}

		pwdService := auth.NewPasswordService(config)

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			result, err := pwdService.VerifyPassword(testCase.password, testCase.hash)

			assert.Equal(t, testCase.result, result, "different results")
			if testCase.err == nil {
				assert.NoError(t, err, "no error expected")
			} else {
				assert.ErrorAs(t, err, &testCase.err, "different errors")
			}
		})
	}
}

func TestPasswordServiceGeneratePasswordHash(t *testing.T) {
	t.Parallel()

	type test struct {
		encryptRepo auth.IEncryptRepo
		password    string
		result      string
		err         error
	}

	testCases := map[string]test{
		"success": {
			encryptRepo: &MockEncryptRepo{
				GenerateResult: "1a2b3c4d",
			},
			password: "password",
			result:   "1a2b3c4d",
		},
		"really long password": {
			encryptRepo: &MockEncryptRepo{
				GenerateResult: "1a2b3c4d",
			},
			password: "this is a really long password, how are you today? i'm doing fine, thanks for asking.",
			result:   "1a2b3c4d",
		},
		"error": {
			encryptRepo: &MockEncryptRepo{
				GenerateErr: errors.New("some hash error"),
			},
			password: "password",
			err:      errors.New("some hash error"),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		config := auth.PasswordConfig{
			EncryptRepo:    testCase.encryptRepo,
			MinLength:      8,
			MaxLength:      20,
			RequireUpper:   true,
			RequireLower:   true,
			RequireNumber:  true,
			RequireSpecial: true,
		}

		pwdService := auth.NewPasswordService(config)

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			result, err := pwdService.GeneratePasswordHash(testCase.password)

			assert.Equal(t, testCase.result, result, "different results")
			if testCase.err == nil {
				assert.NoError(t, err, "no error expected")
			} else {
				assert.ErrorAs(t, err, &testCase.err, "different errors")
			}
		})
	}
}
