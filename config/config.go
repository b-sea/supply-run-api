// Package config defines and parses configurations.
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	JWT struct {
		SignMethod     string `yaml:"signMethod"`
		PublicKeyPath  string `yaml:"publicKeyPath"`
		PrivateKeyPath string `yaml:"privateKeyPath"`
		Issuer         string `yaml:"issuer"`
		Audience       string `yaml:"audience"`
		AccessTimeout  int    `yaml:"accessTimeout"`
		RefreshTimeout int    `yaml:"refreshTimeout"`
	} `yaml:"jwt"`

	Auth struct {
		MinLength int    `yaml:"minLength"`
		MaxLength int    `yaml:"maxLength"`
		Upper     bool   `yaml:"upper"`
		Lower     bool   `yaml:"lower"`
		Number    bool   `yaml:"number"`
		Special   bool   `yaml:"special"`
		Pepper    string `yaml:"pepper"`
	} `yaml:"auth"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load("../../.env"); err != nil {
		logrus.Warnf("error loading .env file: %v", err)
	}

	file, err := os.ReadFile(os.Getenv("SUPPLY_RUN_CONFIG_FILE"))
	if err != nil {
		return nil, fmt.Errorf("unable to read config: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("unable to load config: %v", err)
	}

	return &config, nil
}
