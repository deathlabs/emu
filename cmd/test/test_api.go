package test

import (
	"fmt"
	"net/http"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
)

var (
	testAPICmd = &cobra.Command{
		Use:   "api",
		Short: "Test connectivity to the eMASS API",
		RunE:  testAPI,
	}
)

func testAPI(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		response *http.Response
		system   models.System
		systems  []models.System
	)

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api", config.Data.URL)
		response, err = emass.Get(system.ConfigProfile, endpoint)
		if err != nil {
			return err
		}

		err = output.Response(response, config.OutputFormat)
		if err != nil {
			return err
		}
	}

	return nil
}
