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

	"github.com/deathlabs/emu/models"
	"github.com/deathlabs/emu/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Print EMU configuration information",
		Run:   printConfig,
	}
)

func loadConfig() error {
	var (
		config models.Config
		err    error
		home   string
	)

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err = os.UserHomeDir()
		if err != nil {
			return err
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".emu")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("EMASS_API_KEY")
	viper.AutomaticEnv()
	viper.BindEnv("emass.profiles.production.apiKey", "EMASS_API_KEY_PRODUCTION")
	viper.BindEnv("emass.profiles.pilot.apiKey", "EMASS_API_KEY_PILOT")

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	for i := range config.Systems {
		profile, _ := config.FindProfile(config.Systems[i].ProfileName)
		config.Systems[i].Profile = *profile
	}

	err = config.Validate()
	if err != nil {
		return err
	}

	return nil
}

func printConfig(cmd *cobra.Command, args []string) {
	var (
		config models.Config
		err    error
	)

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output.Config(config, outputFormat)
}

func init() {
	configCmd.AddCommand(&cobra.Command{
		Use:   "config",
		Short: "Print EMU configuration information",
		Run:   printConfig,
	})
	rootCmd.AddCommand(configCmd)
}
