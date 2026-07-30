package get

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
)

const (
	ppsPageIndexDefault = 0
	ppsPageSizeDefault  = 20000
)

var (
	ppsPageIndex int
	ppsPageSize  int
)

var (
	getPpsCmd = &cobra.Command{
		Use:   "pps",
		Short: "Get data about ports, protocols, and services",
		RunE:  getPps,
	}
)

func getPps(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		params   url.Values
		response *http.Response
		system   models.System
		systems  []models.System
	)

	params = url.Values{}

	if ppsPageIndex != ppsPageIndexDefault {
		params.Set("pageIndex", strconv.Itoa(ppsPageIndex))
	}

	if ppsPageSize != ppsPageSizeDefault {
		if ppsPageSize > 20000 {
			return fmt.Errorf("the Page Size cannot exceed %d", ppsPageIndexDefault)
		}
		params.Set("pageSize", strconv.Itoa(ppsPageSize))
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/ports-protocols", config.Data.URL, system.ID)

		if len(params) > 0 {
			endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
		}

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

func init() {
	getPpsCmd.PersistentFlags().IntVarP(&ppsPageIndex, "page-index", "", ppsPageIndexDefault, "Page index")
	getPpsCmd.PersistentFlags().IntVarP(&ppsPageIndex, "page-size", "", ppsPageSizeDefault, "Page size")
}
