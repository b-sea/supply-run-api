package main

func main() {
	// cfg, err := config.Load()
	// if err != nil {
	// 	jsonErr, _ := json.MarshalIndent(errToHTTP(err), "", "    ")
	// 	print(string(jsonErr) + "\n")
	// 	logrus.Fatal("")
	// }

	// pwdConfig := auth.PasswordConfig{
	// 	EncryptRepo:    auth.NewArgon2Repo(auth.Argon2Config{}),
	// 	MinLength:      cfg.Passwords.MinLength,
	// 	MaxLength:      cfg.Passwords.MaxLength,
	// 	RequireUpper:   cfg.Passwords.Upper,
	// 	RequireLower:   cfg.Passwords.Lower,
	// 	RequireNumber:  cfg.Passwords.Number,
	// 	RequireSpecial: cfg.Passwords.Special,
	// }
	// pwdService := auth.NewPasswordService(pwdConfig)

	// publicKey, err := os.ReadFile(cfg.Tokens.PublicKeyPath)
	// if err != nil {
	// 	jsonErr, _ := json.MarshalIndent(errToHTTP(err), "", "    ")
	// 	print(string(jsonErr) + "\n")
	// 	logrus.Fatal("")
	// }
	// privateKey, err := os.ReadFile(cfg.Tokens.PrivateKeyPath)
	// if err != nil {
	// 	jsonErr, _ := json.MarshalIndent(errToHTTP(err), "", "    ")
	// 	print(string(jsonErr) + "\n")
	// 	logrus.Fatal("")
	// }

	// tokenConfig := auth.TokenConfig{
	// 	SignMethod:     cfg.Tokens.SignMethod,
	// 	PublicKey:      publicKey,
	// 	PrivateKey:     privateKey,
	// 	Issuer:         cfg.Tokens.Issuer,
	// 	Audience:       cfg.Tokens.Audience,
	// 	AccessTimeout:  time.Duration(cfg.Tokens.AccessTimeout) * time.Second,
	// 	RefreshTimeout: time.Duration(cfg.Tokens.RefreshTimeout) * time.Second,
	// }
	// tokenService, err := auth.NewTokenService(tokenConfig)
	// if err != nil {
	// 	jsonErr, _ := json.MarshalIndent(errToHTTP(err), "", "    ")
	// 	print(string(jsonErr) + "\n")
	// 	logrus.Fatal("")
	// }

	// accountConfig := service.AccountConfig{
	// 	Password: pwdService,
	// 	Token:    tokenService,
	// 	Repo:     local.NewAccountRepo(),
	// }
	// accountService := service.NewAccountService(accountConfig)

	// seedAccounts := []model.CreateAccountInput{
	// 	{
	// 		Email:    "test-email@email.com",
	// 		Password: "P@ssw0rd1234",
	// 	},
	// 	{
	// 		Email:    "some@email.com",
	// 		Password: "MyPa$$wordswwgjrgew",
	// 	},
	// 	{
	// 		Email:    "bad-email.com.rx",
	// 		Password: "blahblahblahblahblhab;ljsr#$TG2",
	// 	},
	// }

	// ids := []model.ID{}

	// for i := range seedAccounts {
	// 	input := seedAccounts[i]
	// 	result, err := accountService.Signup(input)
	// 	if err != nil {
	// 		logrus.Error(err)
	// 		continue
	// 	}
	// 	ids = append(ids, *result)
	// }

	// i := rand.Intn(len(ids))
	// login, err := accountService.Login(seedAccounts[i].Email, seedAccounts[i].Password)
	// if err != nil {
	// 	jsonErr, _ := json.MarshalIndent(errToHTTP(err), "", "    ")
	// 	print(string(jsonErr) + "\n")
	// 	logrus.Fatal("")
	// }

	// loginJson, _ := json.MarshalIndent(login, "", "    ")
	// print(string(loginJson) + "\n")

	// accessToken, _ := tokenService.ParseAccessToken(login.AccessToken)
	// subject, _ := accessToken.Claims.GetSubject()

	// ctx := accountService.NewContext(context.Background(), subject)
	// account, err := accountService.Profile(ctx)
	// if err != nil {
	// 	jsonErr, _ := json.MarshalIndent(errToHTTP(err), "", "    ")
	// 	print(string(jsonErr) + "\n")
	// 	logrus.Fatal(err)
	// }

	// accountJson, _ := json.MarshalIndent(account, "", "    ")
	// print(string(accountJson) + "\n")

	// time.Sleep(1 * time.Second)

	// newPwd := "bla1Plbahablksag;kwjnsdlgkjshg"
	// err = accountService.Update(ctx, model.UpdateAccountInput{Password: &newPwd})
	// if err != nil {
	// 	jsonErr, _ := json.MarshalIndent(errToHTTP(err), "", "    ")
	// 	print(string(jsonErr) + "\n")
	// 	logrus.Fatal(err)
	// }
	// account, err = accountService.Profile(ctx)
	// if err != nil {
	// 	jsonErr, _ := json.MarshalIndent(errToHTTP(err), "", "    ")
	// 	print(string(jsonErr) + "\n")
	// 	logrus.Fatal(err)
	// }

	// accountJson, _ = json.MarshalIndent(account, "", "    ")
	// print(string(accountJson) + "\n")
}

// type HTTPError struct {
// 	Code    int `json:"-"`
// 	Data    any `json:"data,omitempty"`
// 	message string
// }

// func (e HTTPError) GetMessage() string {
// 	if len(e.message) > 0 {
// 		return e.message
// 	}
// 	return http.StatusText(e.Code)
// }

// func errToHTTP(err error) HTTPError {
// 	// if errors.Is(err, service.ErrAuthentication) {
// 	// 	return HTTPError{
// 	// 		Code: http.StatusUnauthorized,
// 	// 	}
// 	// }

// 	// if errors.Is(err, service.ErrAuthorization) {
// 	// 	return HTTPError{
// 	// 		Code: http.StatusForbidden,
// 	// 	}
// 	// }

// 	// var target auth.InvalidPasswordError
// 	// if errors.As(err, &target) {
// 	// 	return HTTPError{
// 	// 		Code: http.StatusBadRequest,
// 	// 		Data: target.Issues,
// 	// 	}
// 	// }

// 	return HTTPError{
// 		Code: http.StatusInternalServerError,
// 		Data: err.Error(),
// 	}
// }
