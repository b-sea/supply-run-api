package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const envPrefix = "SUPPLY_RUN"

type Config struct {
	Couchbase struct {
		URL  string `yaml:"url"`
		User string `yaml:"user"`
		Pwd  string `yaml:"pwd"`
	} `yaml:"couchbase"`

	Passwords struct {
		MinLen    int  `yaml:"minLen"`
		MaxLen    int  `yaml:"minLen"`
		Uppercase bool `yaml:"uppercase"`
		Lowercase bool `yaml:"lowercase"`
		Number    bool `yaml:"number"`
		Special   bool `yaml:"special"`
	} `yaml:"passwords"`

	Tokens struct {
		SignMethod     string `yaml:"signMethod"`
		PublicKeyPath  string `yaml:"publicKeyPath"`
		PrivateKeyPath string `yaml:"privateKeyPath"`
		Issuer         string `yaml:"issuer"`
		AccessTimeout  int    `yaml:"accessTimeout"`
		RefreshTimeout int    `yaml:"refreshTimeout"`
	} `yaml:"tokens"`

	Argon2 struct {
		Memory     uint32 `yaml:"memory"`
		Passes     uint32 `yaml:"passes"`
		Threads    uint8  `yaml:"threads"`
		SaltLength uint32 `yaml:"saltLength"`
		KeyLength  uint32 `yaml:"keyLength"`
	} `yaml:"argon2"`
}

func Load() (*Config, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		logrus.Debugf("error loading .env file: %s", err)
	}

	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.SetDefault("configFile", "./config.yaml")

	viper.AutomaticEnv()

	logrus.Infof("loading config file: %s", viper.GetString("configFile"))
	viper.SetConfigFile(viper.GetString("configFile"))
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
	return &config, nil
}
