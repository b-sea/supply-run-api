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
}

func Load() (*Config, error) {
	err := godotenv.Load(".env", "../.env")
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
