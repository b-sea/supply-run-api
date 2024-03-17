package auth_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/b-sea/supply-run-api/pkg/auth"
	"github.com/stretchr/testify/assert"
)

const (
	publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCjxzL9vGcGkvp5Mcz2FzmT6wCi
KhLZJlWE5tKAJDiLT1xRNiODNfmAISOwfnmpTJTBqS+92uExymrBY904KSvhYI0o
YqSXdMMvo2AL+nh0+AAuzE9F9nMnCO8NyRtrT6mqqsi762DPiugRleXQmBQaGZFc
hpxqJK75Ybt6H4YQ/QIDAQAB
-----END PUBLIC KEY-----`

	privateKey = `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAKPHMv28ZwaS+nkx
zPYXOZPrAKIqEtkmVYTm0oAkOItPXFE2I4M1+YAhI7B+ealMlMGpL73a4THKasFj
3TgpK+FgjShipJd0wy+jYAv6eHT4AC7MT0X2cycI7w3JG2tPqaqqyLvrYM+K6BGV
5dCYFBoZkVyGnGokrvlhu3ofhhD9AgMBAAECgYAMiZ5MsFSOu0ezaW2QVTzSJwZY
Y7Inr3iSgP0SVDOD7qJZkTRswgAEoATUaQo1PKiSnroJ5ayhnCZKAbQYrYYvTk3U
zSPJq2xgIy7gy1DkAqtHLUqbvVsbbhj0YbyxVYkX0B8MBwD2OSwl1vj6Y21sEOZm
bkIXegoogBZcq1kZXQJBANUTsDl4ymTFsn7pCaYNROHC7fcG4p4IpHJx2jmtZBHT
9w7AorE/qh+RVoJuvDOr5MhTen1HPWNLRHWxqKW6Pk8CQQDExS/1tNEiQWBi0YgC
Ylf2Bt3qKd9khGQ3dUSRXXDn4mOw6JuL0s8TWaN+4qAH9OoJbHQKbbtrOW3hbo6Y
RVTzAkBxOyw77mfH05N+g8Kf3o9LVfZ1ftAw4TDarIwmeHEkVDsHOPF8NfPnIKoT
WFtlLiS/HDWMm64QtS/lR4ryvx1bAkAi+nKWGPh8QGbj6h9lXRoJ0Bquv5bIhYhT
G3N+679gWSwjjJXp+yV4aRzZN2v/PhhEaJUQLYV9gA36Xu7WPqzlAkBlXdjt4hs7
S5j3PfMhRYevcP55+m2x17+x5EQP3XFf9wR3X2U/nNikdvgSedaIavy49uWzEVOp
BcnvQNbMMxiI
-----END PRIVATE KEY-----`

	validAccessToken   = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVC1BQ0NFU1MifQ.eyJpc3MiOiJ1bml0LXRlc3RzIiwic3ViIjoidXNlci1pZCIsImV4cCI6MTA0ODEyNDYyMzYsImlhdCI6MTI1Nzg3NDIwMCwianRpIjoiMTIzNC1teS1pZC01Njc4In0.XGJKFHZiPH_Oqoo11sdc984ZgMQy9hdDCdfSPSmmc8tXnAXeWufYrzui69H8UNlhfGu_becCCefN4xQJKrFo66fWY3hz0MrMAPju-OkiMYXg9TB0_yQTj-APATmemytqcxRFz7l-kitdH8oPGhC5xhf1hv46z6xwQsTVSuX-QHQ`    //nolint: lll
	unknownAccessToken = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVC1BQ0NFU1MifQ.eyJpc3MiOiJzb21lb25lLWVsc2UiLCJzdWIiOiJ1c2VyLWlkIiwiZXhwIjoxMDQ4MTI0NjIzNiwiaWF0IjoxMjU3ODc0MjAwLCJqdGkiOiIxMjM0LW15LWlkLTU2NzgifQ.DXDtA7sxPVhxuS5NZB4E3prxs9V3yyS_u8MqLe2rmXkRf30GNwS9pM0eoMfI0hdNJ7XtnBS91pajhwUL0aZTOnFP7UYf3O49dCFyCYOaw5rixZLxc8AKDqha88BdzU2or33DMPw3KpBjYHAUCV6ykgG7i9SHKkZ8j7AYTDjhVdo` //nolint: lll
	expiredAccessToken = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVC1BQ0NFU1MifQ.eyJpc3MiOiJ1bml0LXRlc3RzIiwic3ViIjoidXNlci1pZCIsImV4cCI6MTI1Nzg3NzgwMCwiaWF0IjoxMjU3ODc0MjAwLCJqdGkiOiIxMjM0LW15LWlkLTU2NzgifQ.BaVcci0hAztvsQQBAsYDVqMjdFSu5ritC0GT82WNNL4q_Vu2Uy3nHSpENL8qCsCm38THnfammGhmKipMlcBTkfy9wIhXlB4uxJ3XG5VfosmFlSDZ5WfdtPgo6kDvmzR_v-JG8q-J7tlhtWhxrrLLib2Lybh7ov6IYQi0-gZ0hEc`     //nolint: lll

	validRefreshToken   = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVC1SRUZSRVNIIn0.eyJpc3MiOiJ1bml0LXRlc3RzIiwic3ViIjoidXNlci1pZCIsImV4cCI6MTA0ODEyNDYyMzYsImlhdCI6MTI1Nzg3NDIwMCwianRpIjoiMTIzNC1teS1pZC01Njc4In0.PkdtDiSiMeaW_4-xo1jZxx1cv4pI2zg4nYrwGrxMx7uNuCMdn7JAYiG9_pi9llJqw8S-1aBFNLf05rkCzXKE3MGx1kRjzPaMDorwETFqVm4JaHOzZJHlfBzryIOzW_RsKMxHFCwiTfxVfahQraesikni2u9BzOfz3e6Msr_3Z3E` //nolint: lll
	unknownRefreshToken = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVC1SRUZSRVNIIn0.eyJpc3MiOiJ1bmtub3duIiwic3ViIjoidXNlci1pZCIsImV4cCI6MTI1Nzg3NzgwMCwiaWF0IjoxMjU3ODc0MjAwLCJqdGkiOiIxMjM0LW15LWlkLTU2NzgifQ.WHaqLsCIUbXDGp_iTt0MqwiWduBWpScvGzkKW3aPDiCOy2UzGP5PedAN0Gk-pDww5rt06t_ejgjwRjK1xFYK7Fpkccsq39I8whLGPeCT9g3JgAmvTJgpB9LYNGD6UBhlrgTd6rN5T0REv_nfLe1bArfxGKUmwZtQC9KP9kPZv2c`      //nolint: lll
	expiredRefreshToken = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVC1SRUZSRVNIIn0.eyJpc3MiOiJ1bml0LXRlc3RzIiwic3ViIjoidXNlci1pZCIsImV4cCI6MTI1Nzg3NzgwMCwiaWF0IjoxMjU3ODc0MjAwLCJqdGkiOiIxMjM0LW15LWlkLTU2NzgifQ.SB6FSk8DPW-PzFDFOtEggDMTa25UbGHW1eaMvg29s3B5CseP4Z5vAdSrcY0Rku35BG0M9zpdSuCNafhTidrOHKzJzDjcmoUjuIrxz6qO_Hzzq0mVk94eBfBhdhLyshK6VZfrjp1KmpiKHA_BpRsfBlZj9yn8FPQkmcRcdmJ-NXM`  //nolint: lll
	badTokenType        = `eyJhbGciOiJSUzI1NiIsInR5cCI6IlJBTkRPTSJ9.eyJpc3MiOiJ1bml0LXRlc3RzIiwic3ViIjoidXNlci1pZCIsImV4cCI6MTA0ODEyNDYyMzYsImlhdCI6MTI1Nzg3NDIwMCwianRpIjoiMTIzNC1teS1pZC01Njc4In0.LN0-BPmxZbIt6KMFsDeJV27fl56MGkRNyNOZ99d4ZbLh9n3H396qzleuE9t0W6EvoMPShXTVMuP3lNc2xL-lJSmynFi2Zk0Odf3nYi0ZRaIc_Txu8OKPYKB2JQUFg2da9QKGv5Wjz50Vx8i3wyLO_DdPjXJtR2XuVk_RObVGBnM`        //nolint: lll
)

func TestTokenService(t *testing.T) {
	t.Parallel()

	type test struct {
		publicKey  []byte
		privateKey []byte
		signMethod string
		err        error
	}

	testCases := map[string]test{
		"success": {
			publicKey:  []byte(publicKey),
			privateKey: []byte(privateKey),
			signMethod: "RS256",
		},
		"unknown sign method": {
			publicKey:  []byte(publicKey),
			privateKey: []byte(privateKey),
			signMethod: "SOMETHING_ELSE",
			err:        auth.ErrRSAKey,
		},
		"bad public key": {
			publicKey:  []byte("-----BAD PUBLIC KEY-----"),
			privateKey: []byte(privateKey),
			signMethod: "RS256",
			err:        auth.ErrRSAKey,
		},
		"bad private key": {
			publicKey:  []byte(publicKey),
			privateKey: []byte("-----BAD PRIVATE KEY-----"),
			signMethod: "RS256",
			err:        auth.ErrRSAKey,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			config := auth.TokenConfig{
				PublicKey:      testCase.publicKey,
				PrivateKey:     testCase.privateKey,
				SignMethod:     testCase.signMethod,
				Issuer:         "unit-tests",
				AccessTimeout:  time.Hour,
				RefreshTimeout: time.Hour,
			}

			_, err := auth.NewTokenService(config)

			if testCase.err == nil {
				assert.NoError(t, err, "no error expected")
			} else {
				assert.ErrorIs(t, err, testCase.err, "different errors")
			}
		})
	}
}

func TestTokenServiceGenerateAccessToken(t *testing.T) {
	t.Parallel()

	auth.Timestamp = func() time.Time {
		return time.Date(2009, 11, 10, 12, 30, 0, 0, time.Local)
	}

	type test struct {
		sub      string
		issuer   string
		audience string
		result   string
		err      error
	}

	testCases := map[string]test{
		"success": {
			sub:      "user-id",
			issuer:   "unit-tests",
			audience: "special-service",
			result: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1bml0LXRlc3RzIiwic3ViIjoidXNlci1pZCIsImF1ZCI6WyJh" +
				"Y2Nlc3MiLCJzcGVjaWFsLXNlcnZpY2UiXSwiZXhwIjoxMjU3ODc3ODAwLCJpYXQiOjEyNTc4NzQyMDAsImp0aSI6IjEyMzQtbXkta" +
				"WQtNTY3OCJ9.CN_SviUPwxugvy80zfkD8v4REQjteaqgN4kxrBPOgF0-14rn19MFkPdgaBRX2B-t3CNR9zP9MwzLvEdmZwGGDwrhP" +
				"1K6MCQp1CKwtRR02XufSNjtp94_jJR6cUHQot5EbpVq7uKqeWKni-nGImfiiwQjP-MQZd9bBjs9ZeikTSc",
		},
		"no issuer": {
			sub:      "user-id",
			audience: "special-service",
			result: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ1c2VyLWlkIiwiYXVkIjpbImFjY2VzcyIsInNwZWNpYWwtc2Vy" +
				"dmljZSJdLCJleHAiOjEyNTc4Nzc4MDAsImlhdCI6MTI1Nzg3NDIwMCwianRpIjoiMTIzNC1teS1pZC01Njc4In0.nqcFQpMT7vN3E" +
				"WlWdVGkZLnhAheAVd2EbYfLMltOM3K8CRCankAfepFpluSCX4KvMVoez5UgjMa0aHuKV9b-M3S2D8gRp0R6akrd8AAZjqEoLy0dhb" +
				"BQoLvt1sYENzTtgcvi8Qt1rKi_WY9wm5Awjz3TT-PUplmzEzIRGbi6Ctw",
		},
		"no audience": {
			sub:    "user-id",
			issuer: "unit-tests",
			result: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1bml0LXRlc3RzIiwic3ViIjoidXNlci1pZCIsImF1ZCI6WyJh" +
				"Y2Nlc3MiXSwiZXhwIjoxMjU3ODc3ODAwLCJpYXQiOjEyNTc4NzQyMDAsImp0aSI6IjEyMzQtbXktaWQtNTY3OCJ9.SzhzvwzUsxtD" +
				"EWj_hM-NoM9mFdTxN5PJp6d8VdGxKam6J6Tm4vxhhYJBfl-pd_HKWcERheu4R4fFtTD0bIdxNK1OqW-vwCecBojrPavuzVM7PhWd8" +
				"ne5QMImDpnSyxo7iRSbmV_uIi5EiPIa5KPtECNmAXOACqEidsLY9IILlaQ",
		},
		"no subject": {
			issuer:   "unit-tests",
			audience: "special-service",
			err:      auth.ErrJWTClaim,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		config := auth.TokenConfig{
			PublicKey:      []byte(publicKey),
			PrivateKey:     []byte(privateKey),
			SignMethod:     "RS256",
			Issuer:         testCase.issuer,
			Audience:       testCase.audience,
			AccessTimeout:  time.Hour,
			RefreshTimeout: time.Hour,
			IDGenerator: func() string {
				return "1234-my-id-5678"
			},
		}

		tokenService, err := auth.NewTokenService(config)
		if err == nil {
			assert.NoError(t, err, "no error expected during service creation")
		}

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			result, err := tokenService.GenerateAccessToken(testCase.sub)

			assert.Equal(t, testCase.result, result, "different results")
			if testCase.err == nil {
				assert.NoError(t, err, "no error expected")
			} else {
				assert.ErrorIs(t, err, testCase.err, "different errors")
			}
		})
	}
}

func TestTokenServiceGenerateRefreshToken(t *testing.T) {
	t.Parallel()

	auth.Timestamp = func() time.Time {
		return time.Date(2009, 11, 10, 12, 30, 0, 0, time.Local)
	}

	type test struct {
		sub      string
		issuer   string
		audience string
		result   string
		err      error
	}

	testCases := map[string]test{
		"success": {
			sub:      "user-id",
			issuer:   "unit-tests",
			audience: "special-service",
			result: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1bml0LXRlc3RzIiwic3ViIjoidXNlci1pZCIsImF1ZCI6WyJh" +
				"Y2Nlc3MiLCJzcGVjaWFsLXNlcnZpY2UiXSwiZXhwIjoxMjU3ODc3ODAwLCJpYXQiOjEyNTc4NzQyMDAsImp0aSI6IjEyMzQtbXkta" +
				"WQtNTY3OCJ9.CN_SviUPwxugvy80zfkD8v4REQjteaqgN4kxrBPOgF0-14rn19MFkPdgaBRX2B-t3CNR9zP9MwzLvEdmZwGGDwrhP" +
				"1K6MCQp1CKwtRR02XufSNjtp94_jJR6cUHQot5EbpVq7uKqeWKni-nGImfiiwQjP-MQZd9bBjs9ZeikTSc",
		},
		"no issuer": {
			sub:      "user-id",
			audience: "special-service",
			result: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ1c2VyLWlkIiwiYXVkIjpbImFjY2VzcyIsInNwZWNpYWwtc2Vy" +
				"dmljZSJdLCJleHAiOjEyNTc4Nzc4MDAsImlhdCI6MTI1Nzg3NDIwMCwianRpIjoiMTIzNC1teS1pZC01Njc4In0.nqcFQpMT7vN3E" +
				"WlWdVGkZLnhAheAVd2EbYfLMltOM3K8CRCankAfepFpluSCX4KvMVoez5UgjMa0aHuKV9b-M3S2D8gRp0R6akrd8AAZjqEoLy0dhb" +
				"BQoLvt1sYENzTtgcvi8Qt1rKi_WY9wm5Awjz3TT-PUplmzEzIRGbi6Ctw",
		},
		"no audience": {
			sub:    "user-id",
			issuer: "unit-tests",
			result: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1bml0LXRlc3RzIiwic3ViIjoidXNlci1pZCIsImF1ZCI6WyJh" +
				"Y2Nlc3MiXSwiZXhwIjoxMjU3ODc3ODAwLCJpYXQiOjEyNTc4NzQyMDAsImp0aSI6IjEyMzQtbXktaWQtNTY3OCJ9.SzhzvwzUsxtD" +
				"EWj_hM-NoM9mFdTxN5PJp6d8VdGxKam6J6Tm4vxhhYJBfl-pd_HKWcERheu4R4fFtTD0bIdxNK1OqW-vwCecBojrPavuzVM7PhWd8" +
				"ne5QMImDpnSyxo7iRSbmV_uIi5EiPIa5KPtECNmAXOACqEidsLY9IILlaQ",
		},
		"no subject": {
			issuer:   "unit-tests",
			audience: "special-service",
			err:      auth.ErrJWTClaim,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		config := auth.TokenConfig{
			PublicKey:      []byte(publicKey),
			PrivateKey:     []byte(privateKey),
			SignMethod:     "RS256",
			Issuer:         testCase.issuer,
			Audience:       testCase.audience,
			AccessTimeout:  time.Hour,
			RefreshTimeout: time.Hour,
			IDGenerator: func() string {
				return "1234-my-id-5678"
			},
		}

		tokenService, err := auth.NewTokenService(config)
		if err == nil {
			assert.NoError(t, err, "no error expected during service creation")
		}

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			result, err := tokenService.GenerateRefreshToken(testCase.sub)

			assert.Equal(t, testCase.result, result, "different results")
			if testCase.err == nil {
				assert.NoError(t, err, "no error expected")
			} else {
				assert.ErrorIs(t, err, testCase.err, "different errors")
			}
		})
	}
}

// func TestTokenServiceParseAccessToken(t *testing.T) {
// 	t.Parallel()

// 	type test struct {
// 		token  string
// 		result *jwt.Token
// 		err    error
// 	}

// 	testCases := map[string]test{
// 		"success": {
// 			token: validAccessToken,
// 			result: &jwt.Token{
// 				Raw:    validAccessToken,
// 				Method: jwt.SigningMethodRS256,
// 				Header: map[string]interface{}{
// 					"typ": "JWT-ACCESS",
// 					"alg": jwt.SigningMethodRS256.Alg(),
// 				},
// 				Claims: &jwt.RegisteredClaims{
// 					ID:        "1234-my-id-5678",
// 					Issuer:    "unit-tests",
// 					Subject:   "user-id",
// 					ExpiresAt: jwt.NewNumericDate(time.Unix(int64(10481246236), 0)),
// 					IssuedAt:  jwt.NewNumericDate(time.Unix(int64(1257874200), 0)),
// 				},
// 				Valid: true,
// 				Signature: []byte(
// 					"\\bJ\x14vb<\x7fΪ\x8a5\xd6\xc7\\\xf7\xce\x19\x80\xc42\xf6\x17C\t\xd7\xd2=)\xa6s\xcbW" +
// 						"\x9c\x05\xdeZ\xe7د;\xa2\xeb\xd1\xfcP\xd9a|k\xbfm\xe7\x02\t\xe7\xcd\xe3\x14\t*\xb1h" +
// 						"\xeb\xa7\xd6cxs\xd0\xca\xcc\x00\xf8\xee\xf8\xe9\"1\x85\xe0\xf50t\xff$\x13\x8f\xe0" +
// 						"\x0f\x019\x9e\x9b+js\x14EϹ~\x92+]\x1f\xca\x0f\x1a\x10\xb9\xc6\x17\xf5\x86\xfe:ϬpB" +
// 						"\xc4\xd5J\xe5\xfe@t",
// 				),
// 			},
// 		},
// 		"bad type": {
// 			token: badTokenType,
// 			err:   jwt.ErrTokenUnverifiable,
// 		},
// 		"bad issuer": {
// 			token: unknownAccessToken,
// 			err:   jwt.ErrTokenInvalidIssuer,
// 		},
// 		"expired": {
// 			token: expiredAccessToken,
// 			err:   jwt.ErrTokenExpired,
// 		},
// 	}

// 	for name, testCase := range testCases {
// 		name, testCase := name, testCase

// 		config := auth.TokenConfig{
// 			PublicKey:      []byte(publicKey),
// 			PrivateKey:     []byte(privateKey),
// 			SignMethod:     "RS256",
// 			Issuer:         "unit-tests",
// 			AccessTimeout:  time.Hour,
// 			RefreshTimeout: time.Hour,
// 		}

// 		tokenService, err := auth.NewTokenService(config)
// 		if err == nil {
// 			assert.NoError(t, err, "no error expected during service creation")
// 		}

// 		t.Run(name, func(s *testing.T) {
// 			s.Parallel()

// 			result, err := tokenService.ParseAccessToken(testCase.token)

// 			assert.Equal(t, testCase.result, result, "different results")
// 			if testCase.err == nil {
// 				assert.NoError(t, err, "no error expected")
// 			} else {
// 				assert.ErrorIs(t, err, testCase.err, "different errors")
// 			}
// 		})
// 	}
// }

// func TestTokenServiceParseRefreshToken(t *testing.T) {
// 	t.Parallel()

// 	type test struct {
// 		token  string
// 		result *jwt.Token
// 		err    error
// 	}

// 	testCases := map[string]test{
// 		"success": {
// 			token: validRefreshToken,
// 			result: &jwt.Token{
// 				Raw:    validRefreshToken,
// 				Method: jwt.SigningMethodRS256,
// 				Header: map[string]interface{}{
// 					"typ": "JWT-REFRESH",
// 					"alg": jwt.SigningMethodRS256.Alg(),
// 				},
// 				Claims: &jwt.RegisteredClaims{
// 					ID:        "1234-my-id-5678",
// 					Issuer:    "unit-tests",
// 					Subject:   "user-id",
// 					ExpiresAt: jwt.NewNumericDate(time.Unix(int64(10481246236), 0)),
// 					IssuedAt:  jwt.NewNumericDate(time.Unix(int64(1257874200), 0)),
// 				},
// 				Valid: true,
// 				Signature: []byte(
// 					">Gm\x0e$\xa21\xe6\x96\xff\x8f\xb1\xa3X\xd9\xc7\x1d\\\xbf\x8aH\xdb88\x9d\x8a\xf0\x1a" +
// 						"\xbcLǻ\x8d\xb8#\x1d\x9f\xb2@b!\xbd\xfe\x98\xbd\x96Rj\xc3ľՠE4\xb7\xf4\xe6\xb9\x02" +
// 						"\xcdr\x84\xdc\xc1\xb1\xd6Dc\xcc\xf6\x8c\x0e\x8a\xf0\x111jVn\ths\xb3d\x91\xe5|\x1c" +
// 						"\xebȃ\xb3[\xf4l(\xccG\x14,\"M\xfcU}\xa8P\xad\xa7\xac\x8aI\xe2\xda\xefA\xcc\xe7\xf3" +
// 						"\xdd\ue332\xbf\xf7gq",
// 				),
// 			},
// 		},
// 		"bad type": {
// 			token: badTokenType,
// 			err:   jwt.ErrTokenUnverifiable,
// 		},
// 		"bad issuer": {
// 			token: unknownRefreshToken,
// 			err:   jwt.ErrTokenInvalidIssuer,
// 		},
// 		"expired": {
// 			token: expiredRefreshToken,
// 			err:   jwt.ErrTokenExpired,
// 		},
// 	}

// 	for name, testCase := range testCases {
// 		name, testCase := name, testCase

// 		config := auth.TokenConfig{
// 			PublicKey:      []byte(publicKey),
// 			PrivateKey:     []byte(privateKey),
// 			SignMethod:     "RS256",
// 			Issuer:         "unit-tests",
// 			AccessTimeout:  time.Hour,
// 			RefreshTimeout: time.Hour,
// 		}

// 		tokenService, err := auth.NewTokenService(config)
// 		if err == nil {
// 			assert.NoError(t, err, "no error expected during service creation")
// 		}

// 		t.Run(name, func(s *testing.T) {
// 			s.Parallel()

// 			result, err := tokenService.ParseRefreshToken(testCase.token)

// 			assert.Equal(t, testCase.result, result, "different results")
// 			if testCase.err == nil {
// 				assert.NoError(t, err, "no error expected")
// 			} else {
// 				assert.ErrorIs(t, err, testCase.err, "different errors")
// 			}
// 		})
// 	}
// }

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

		config := auth.TokenConfig{
			PublicKey:      []byte(publicKey),
			PrivateKey:     []byte(privateKey),
			SignMethod:     "RS256",
			Issuer:         "unit-tests",
			AccessTimeout:  time.Hour,
			RefreshTimeout: time.Hour,
		}

		tokenService, err := auth.NewTokenService(config)
		if err == nil {
			assert.NoError(t, err, "no error expected during service creation")
		}

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			result, ok := tokenService.FromHeader(testCase.header)

			assert.Equal(t, testCase.result, result, "different results")
			assert.Equal(t, testCase.ok, ok, "different ok")
		})
	}
}
