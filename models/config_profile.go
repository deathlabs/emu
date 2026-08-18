package models

import (
	"encoding/json"
	"strings"
)

type ConfigProfile struct {
	Name           string `mapstructure:"name"           json:"name"           yaml:"name"`
	PublicKeyPath  string `mapstructure:"publicKeyPath"  json:"publicKeyPath"  yaml:"publicKeyPath"`
	PrivateKeyPath string `mapstructure:"privateKeyPath" json:"privateKeyPath" yaml:"privateKeyPath"`
	APIKey         string `mapstructure:"-"              json:"apiKey"         yaml:"apiKey"`
	UserUID        string `mapstructure:"-"              json:"userUID"        yaml:"userUID"`
}

func (profile ConfigProfile) MaskedAPIKey() string {
	var (
		apiKeyLength       int
		lastFourCharacters string
	)

	if len(profile.APIKey) == 0 {
		return ""
	}

	apiKeyLength = len(profile.APIKey)
	lastFourCharacters = profile.APIKey[apiKeyLength-4:]
	return strings.Repeat("*", apiKeyLength-4) + lastFourCharacters
}

func (profile ConfigProfile) MaskedUserUID() string {
	var (
		userUIDLength      int
		lastFourCharacters string
	)

	if len(profile.UserUID) == 0 {
		return ""
	}

	userUIDLength = len(profile.UserUID)
	lastFourCharacters = profile.UserUID[userUIDLength-4:]
	return strings.Repeat("*", userUIDLength-4) + lastFourCharacters
}

func (profile ConfigProfile) MarshalYAML() (interface{}, error) {
	// Create a new struct that embeds the alias and adds the masked API key field.
	type Output struct {
		Name           string `yaml:"name"`
		PublicKeyPath  string `yaml:"publicKeyPath"`
		PrivateKeyPath string `yaml:"privateKeyPath"`
		APIKey         string `yaml:"apiKey"`
		UserUID        string `yaml:"userUID"`
	}

	// Return the config profile with the masked API key and no errors.
	return Output{
		Name:           profile.Name,
		PublicKeyPath:  profile.PublicKeyPath,
		PrivateKeyPath: profile.PrivateKeyPath,
		APIKey:         profile.MaskedAPIKey(),
		UserUID:        profile.MaskedUserUID(),
	}, nil
}

func (profile ConfigProfile) MarshalJSON() ([]byte, error) {
	// Create an alias to avoid infinite recursion when calling json.Marshal within this method.
	type Alias ConfigProfile

	// Create a new struct that embeds the alias and adds the masked API key field.
	type Output struct {
		Alias
		APIKey  string `json:"apiKey"`
		UserUID string `json:"userUID"`
	}

	// Return the config profile with the masked API key and no errors.
	return json.Marshal(Output{
		Alias:   Alias(profile),
		APIKey:  profile.MaskedAPIKey(),
		UserUID: profile.MaskedUserUID(),
	})
}
