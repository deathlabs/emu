/*
Copyright © 2026 Vic Fernandez III <@cyberphor>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package models

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	URL            string          `mapstructure:"url" json:"url" yaml:"url"`
	ConfigProfiles []ConfigProfile `mapstructure:"profiles" json:"profiles" yaml:"profiles"`
	Systems        []System        `mapstructure:"systems" json:"systems" yaml:"systems"`
	Settings       struct {
		Output struct {
			Format string `mapstructure:"format" json:"format" yaml:"format"`
		} `mapstructure:"output" json:"output" yaml:"output"`
	} `mapstructure:"settings" json:"settings" yaml:"settings"`
}

func (config *Config) ResolveProfilesToSystems() {
	var (
		ok           bool
		profile      ConfigProfile
		profileIndex int
		profiles     = make(map[string]ConfigProfile)
		profileName  string
		system       map[string]interface{}
		systemIndex  int
		systems      []interface{}
	)

	// Set each profile's API key using the corresponding environment variable (i.e., EMASS_API_KEY_<PROFILE_NAME>).
	for _, profile = range config.ConfigProfiles {
		// Get the current's profile name from configuration.
		profileName = strings.ToUpper(profile.Name)

		// Get the API key for the current profile using the corresponding environment variable.
		profile.APIKey = os.Getenv("EMASS_API_KEY_" + profileName)

		// Get the user UID for the current profile using the corresponding environment variable.
		profile.UserUID = os.Getenv("EMASS_USER_UID_" + profileName)

		// Save the updated profile in the profiles map (required to resolve the systems' profiles in the loop below).
		profiles[profile.Name] = profile
	}

	// Save the updated profiles map in the config struct.
	for profileIndex, profile = range config.ConfigProfiles {
		config.ConfigProfiles[profileIndex] = profiles[profile.Name]
	}

	// Get all systems in the config as a slice of interfaces.
	systems = viper.Get("systems").([]interface{})

	// Resolve the profiles of each system based on the profile name specified in the config.
	for systemIndex = range config.Systems {

		// Get the current system as a map.
		system = systems[systemIndex].(map[string]interface{})

		// Get the current system's profile name as a string
		profileName = system["profile"].(string)

		// Lookup the profile name in the profiles map.
		profile, ok = profiles[profileName]
		if ok {
			config.Systems[systemIndex].ConfigProfile = profile
		}
	}
}

func (config *Config) GetProfileBySystemID(systemID int) (ConfigProfile, error) {
	var system System

	for _, system = range config.Systems {
		if system.ID == systemID {
			return system.ConfigProfile, nil
		}
	}

	return ConfigProfile{}, fmt.Errorf("no profile found for system id %d", systemID)
}

func (config *Config) GetProfileByName(profileName string) (ConfigProfile, error) {
	var profile ConfigProfile

	for _, profile = range config.ConfigProfiles {
		if profile.Name == profileName {
			return profile, nil
		}
	}

	return ConfigProfile{}, fmt.Errorf("no profile found for name %s", profileName)
}
