// Package config defines and parses the application configuration.
package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const envPrefix = "SUPPLY_RUN"

// Config defines the configuration values for the application.
type Config struct {
	Couchbase struct {
		URL  string `yaml:"url"`
		User string `yaml:"user"`
		Pwd  string `yaml:"pwd"`
	} `yaml:"couchbase"`

	Tokens struct {
		SignMethod     string `yaml:"signMethod"`
		PublicKeyPath  string `yaml:"publicKeyPath"`
		PrivateKeyPath string `yaml:"privateKeyPath"`
		Issuer         string `yaml:"issuer"`
		Audience       string `yaml:"audience"`
		AccessTimeout  int    `yaml:"accessTimeout"`
		RefreshTimeout int    `yaml:"refreshTimeout"`
	} `yaml:"tokens"`

	Passwords struct {
		Pepper    string `yaml:"pepper"`
		MinLength int    `yaml:"minLength"`
		MaxLength int    `yaml:"maxLength"`
		Upper     bool   `yaml:"upper"`
		Lower     bool   `yaml:"lower"`
		Number    bool   `yaml:"number"`
		Special   bool   `yaml:"special"`
	} `yaml:"passwords"`
}

// Load parses the application configuration.
func Load() (*Config, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		logrus.Warnf("error loading .env file: %s", err)
	}

	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.SetDefault("config_file", "./config.yaml")

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
