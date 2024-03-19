// Package configs defines and parses configurations.
package configs

import (
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Defaults defines all default objects.
type Defaults struct {
	Units []struct {
		ID     uuid.UUID `yaml:"id"`
		Name   string    `yaml:"name"`
		Symbol string    `yaml:"symbol"`
		System string    `yaml:"system"`
		Type   string    `yaml:"type"`
	} `yaml:"units"`
}

// LoadDefaults parses all default objects.
func LoadDefaults() *Defaults {
	if err := godotenv.Load("../../.env"); err != nil {
		logrus.Warnf("error loading .env file: %v", err)
	}

	file, err := os.ReadFile(os.Getenv("SUPPLY_RUN_DEFAULTS_YAML"))
	if err != nil {
		logrus.Errorf("unable to read defaults config: %v", err)
		return nil
	}

	var defaults Defaults
	if err := yaml.Unmarshal(file, &defaults); err != nil {
		logrus.Error("unable to load ", err)
	}

	return &defaults
}
