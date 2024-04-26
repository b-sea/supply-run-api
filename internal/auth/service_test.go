package auth_test

import (
	"net/http"
	"testing"

	"github.com/b-sea/supply-run-api/internal/auth"
	"github.com/stretchr/testify/assert"
)

func TestFromHeader(t *testing.T) {
	t.Parallel()

	type test struct {
		header http.Header
		result string
		ok     bool
	}

	testCases := map[string]test{
		"success": {
			header: http.Header{
				"Authorization": []string{"Bearer --my-token--"},
			},
			result: "--my-token--",
			ok:     true,
		},
		"no header": {
			header: http.Header{},
			ok:     false,
		},
		"multiple headers": {
			header: http.Header{
				"Authorization": []string{
					"Bearer --my-token--",
					"Bearer another-token!",
				},
			},
			ok: false,
		},
		"non bearer": {
			header: http.Header{
				"Authorization": []string{"Basic user:pwd(but in base64)"},
			},
			ok: false,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			result, ok := auth.TokenFromHeader(testCase.header)

			assert.Equal(t, testCase.result, result, "different results")
			assert.Equal(t, testCase.ok, ok, "different ok")
		})
	}
}
