package get

import (
	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
)

var (
	getConfigCmd = &cobra.Command{
		Use:   "config",
		Short: "Print EMU configuration information",
		RunE:  outputConfig,
	}
)

func outputConfig(cmd *cobra.Command, args []string) error {
	var err = output.Config(config.Data, config.OutputFormat)
	if err != nil {
		return err
	}
	return nil
}
