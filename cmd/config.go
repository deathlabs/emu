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
	"path/filepath"

	"github.com/deathlabs/emu/models"
	"github.com/deathlabs/emu/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config    models.Config
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Print EMU configuration information",
		Run:   printConfig,
	}
)

func loadConfig(cmd *cobra.Command, args []string) error {
	var (
		err error
	)

	if configFile != DefaultConfigFilePath {
		viper.SetConfigType("yaml")
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigFile(filepath.Join(".", DefaultConfigFilePath))
	}

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	config.ResolveProfilesToSystems()

	return nil
}

func printConfig(cmd *cobra.Command, args []string) {
	output.Config(config, outputFormat)
}

func init() {
	rootCmd.AddCommand(configCmd)
}
