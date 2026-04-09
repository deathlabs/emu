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
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/deathlabs/emu/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	DefaultConfigFilePath = ".emu.yaml"
	DefaultOutputFormat   = "json"
)

var (
	activeProfileName string
	config            models.Config
	configFile        string
	outputFormat      string
	rootCmd           = &cobra.Command{
		Use:   "emu",
		Short: "eMASS Updater (EMU) is a tool for automating eMASS records management.",
	}
	systemIDs []int
)

func checkArguments() error {
	switch outputFormat {
	case "json", "yaml":
	default:
		return fmt.Errorf("invalid output format \"%s\"", outputFormat)
	}
	return nil
}

func setupEMASSClient(cmd *cobra.Command, args []string) error {
	var err error

	// Check arguments passed to the root command.
	err = checkArguments()
	if err != nil {
		return err
	}

	// Set the config filepath to the default if one is not provided.
	if configFile != DefaultConfigFilePath {
		viper.SetConfigType("yaml")
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigFile(filepath.Join(".", DefaultConfigFilePath))
	}

	// Copy the config into memory.
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	// Unmarshal the config into a Config struct.
	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	// Resolve the profiles for each system in the config.
	config.ResolveProfilesToSystems()

	return nil
}

func filterProfiles(config models.Config, activeProfileName string) ([]models.ConfigProfile, error) {
	var (
		profile  models.ConfigProfile
		profiles []models.ConfigProfile
	)

	if activeProfileName == "" {
		return config.ConfigProfiles, nil
	}

	for _, profile = range config.ConfigProfiles {
		if profile.Name == activeProfileName {
			profiles = append(profiles, profile)
			return profiles, nil
		}
	}

	return nil, fmt.Errorf("no profile found for name %s", activeProfileName)
}

func filterSystems(config models.Config, profileName string, systemIDs []int) ([]models.System, error) {
	var (
		filteredSystems []models.System
		system          models.System
	)

	for _, system = range config.Systems {
		if profileName != "" && system.ConfigProfile.Name != profileName {
			continue
		}

		if len(systemIDs) > 0 && !containsSystemID(systemIDs, system.ID) {
			continue
		}

		filteredSystems = append(filteredSystems, system)
	}

	if len(filteredSystems) == 0 {
		return nil, fmt.Errorf("no systems matched the requested filters")
	}

	return filteredSystems, nil
}

func containsSystemID(ids []int, id int) bool {
	var current int

	for _, current = range ids {
		if current == id {
			return true
		}
	}
	return false
}

func Execute() {
	var err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Register flags.
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", DefaultConfigFilePath, "Config file path")
	rootCmd.PersistentFlags().StringVarP(&activeProfileName, "profile", "p", "", "Config profile name")
	rootCmd.PersistentFlags().IntSliceVarP(&systemIDs, "system-id", "s", []int{}, "System IDs (can specify multiple)")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", DefaultOutputFormat, "Output format (json or yaml)")

	// Setup the eMASS client before executing the root command (i.e., any command).
	rootCmd.PersistentPreRunE = setupEMASSClient
}
