/*
Copyright © 2026 Vic Fernandez <@cyberphor>

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

	"github.com/deathlabs/emu/output"
	"github.com/deathlabs/emu/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Create an EMU config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create config called")
	},
}

func loadConfig() error {
	var (
		err  error
		home string
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

	viper.SetDefault("timeout", 30)
	viper.SetDefault("output_format", outputFormat)
	viper.SetDefault("headers", map[string]string{
		"User-Agent": "emu/v4.0.0",
	})

	viper.SetEnvPrefix("EMASS")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.Get("emass")

	return nil
}

func getConfig(cmd *cobra.Command, args []string) {
	var (
		config types.Config
		err    error
	)
	err = viper.Unmarshal(&config)

	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		os.Exit(1)
	}
	output.Config(config, outputFormat)
}

var getConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Print your EMU config",
	Run:   getConfig,
}

func init() {
	createCmd.AddCommand(createConfigCmd)
	getCmd.AddCommand(getConfigCmd)
}
