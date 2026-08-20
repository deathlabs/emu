package cmd

import (
	"fmt"
	"os"

	"github.com/deathlabs/emu/v4/cmd/create"
	"github.com/deathlabs/emu/v4/cmd/delete"
	"github.com/deathlabs/emu/v4/cmd/get"
	"github.com/deathlabs/emu/v4/cmd/test"
	"github.com/deathlabs/emu/v4/cmd/update"
	"github.com/deathlabs/emu/v4/cmd/upload"
	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/spf13/cobra"
)

const (
	emuVersion   = "v4.0.5"
	emassVersion = "v3.32.0"
)

var (
	rootCmd = &cobra.Command{
		Use:     "emu",
		Short:   "eMASS Updater (EMU) is a tool for automating eMASS records management.",
		Version: fmt.Sprintf("%s\neMASS API version %s", emuVersion, emassVersion),
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
	rootCmd.AddCommand(test.Cmd)
	rootCmd.AddCommand(update.Cmd)
	rootCmd.AddCommand(upload.Cmd)
}
