package models

import (
	"errors"
	"fmt"
)

type Profile struct {
	Name           string `json:"name" yaml:"name"`
	APIKey         string `json:"apiKey" yaml:"apiKey"`
	PublicKeyPath  string `json:"publicKeyPath" yaml:"publicKeyPath"`
	PrivateKeyPath string `json:"privateKeyPath" yaml:"privateKeyPath"`
}

type System struct {
	Name        string  `json:"name" yaml:"name"`
	ID          int     `json:"id" yaml:"id"`
	ProfileName string  `json:"profileName" yaml:"profileName"`
	Profile     Profile `json:"profile" yaml:"profile"`
}

type Config struct {
	URL      string    `json:"url" yaml:"url"`
	Profiles []Profile `json:"profiles" yaml:"profiles"`
	Systems  []System  `json:"systems" yaml:"systems"`
	Settings struct {
		Output struct {
			Format string `json:"format" yaml:"format"`
		} `json:"output" yaml:"output"`
	} `json:"settings" yaml:"settings"`
}

func (config *Config) Validate() error {
	if config.URL == "" {
		return errors.New("a URL to eMASS was not provided")
	}
	return nil
}

func (config *Config) FindProfile(name string) (*Profile, error) {
	for i := range config.Profiles {
		if config.Profiles[i].Name == name {
			return &config.Profiles[i], nil
		}
	}
	return nil, fmt.Errorf("profile %q not found", name)
}
