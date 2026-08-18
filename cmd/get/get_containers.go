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
	containersPageIndexDefault = 0
	containersPageSizeDefault  = 20000
)

var (
	containersPageIndex int
	containersPageSize  int
)

var (
	getContainersCmd = &cobra.Command{
		Use:   "containers",
		Short: "Get data about containers",
		RunE:  getContainers,
	}
)

func getContainers(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		params   url.Values
		response *http.Response
		system   models.System
		systems  []models.System
	)

	params = url.Values{}

	if containersPageIndex != containersPageIndexDefault {
		params.Set("pageIndex", strconv.Itoa(containersPageIndex))
	}

	if containersPageSize != containersPageSizeDefault {
		if containersPageSize > 20000 {
			return fmt.Errorf("the Page Size cannot exceed %d", containersPageIndexDefault)
		}
		params.Set("pageSize", strconv.Itoa(containersPageSize))
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/containers", config.Data.URL, system.ID)

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
	getContainersCmd.PersistentFlags().IntVarP(&containersPageIndex, "page-index", "", containersPageIndexDefault, "Page index")
	getContainersCmd.PersistentFlags().IntVarP(&containersPageSize, "page-size", "", containersPageSizeDefault, "Page size")
}
