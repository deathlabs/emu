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
	"os"

	"github.com/deathlabs/emu/cmd/create"
	"github.com/deathlabs/emu/cmd/delete"
	"github.com/deathlabs/emu/cmd/get"
	"github.com/deathlabs/emu/cmd/update"
	"github.com/deathlabs/emu/cmd/upload"
	"github.com/deathlabs/emu/config"
	"github.com/deathlabs/emu/emass"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "emu",
		Short: "eMASS Updater (EMU) is a tool for automating eMASS records management.",
	}
)

func Execute() {
	var err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Define persistent flags for the root command.
	rootCmd.PersistentFlags().StringVarP(&config.Filename, "config", "c", config.DefaultConfigFilePath, "Config file path")
	rootCmd.PersistentFlags().StringVarP(&config.ActiveProfileName, "profile", "p", "", "Config profile name")
	rootCmd.PersistentFlags().IntSliceVarP(&config.SystemIDs, "system-id", "s", []int{}, "System IDs (can specify multiple)")
	rootCmd.PersistentFlags().StringVarP(&config.OutputFormat, "output", "o", config.DefaultOutputFormat, "Output format (json or yaml)")

	// Setup the eMASS client before executing the root command (i.e., any command).
	rootCmd.PersistentPreRunE = emass.SetupClient

	// Add subcommands to the root command.
	rootCmd.AddCommand(create.Cmd)
	rootCmd.AddCommand(delete.Cmd)
	rootCmd.AddCommand(get.Cmd)
	rootCmd.AddCommand(update.Cmd)
	rootCmd.AddCommand(upload.Cmd)
}
