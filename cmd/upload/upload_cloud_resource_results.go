package upload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v3"
)

var (
	cloudResourceResultsFilePath string
)

var (
	uploadCloudResourceResultsCmd = &cobra.Command{
		Use:   "cloud-resource-results",
		Short: "Upload cloud resource results",
		RunE:  uploadCloudResourceResults,
	}
)

func uploadCloudResourceResults(cmd *cobra.Command, args []string) error {
	var (
		cloudResourceResults     []models.CloudResourceResult
		cloudResourceResultsFile []byte
		cloudResourceResultsJson []byte
		err                      error
		endpoint                 string
		headers                  map[string]string
		request                  *bytes.Buffer
		response                 *http.Response
		system                   models.System
		systems                  []models.System
	)

	headers = map[string]string{
		"Content-Type": "application/json",
	}

	cloudResourceResultsFile, err = os.ReadFile(cloudResourceResultsFilePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(cloudResourceResultsFile, &cloudResourceResults)
	if err != nil {
		return err
	}

	cloudResourceResultsJson, err = json.Marshal(cloudResourceResults)
	if err != nil {
		return err
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/cloud-resource-results", config.Data.URL, system.ID)

		request = bytes.NewBuffer(cloudResourceResultsJson)

		response, err = emass.Post(system.ConfigProfile, endpoint, headers, request)
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
	uploadCloudResourceResultsCmd.PersistentFlags().StringVarP(&cloudResourceResultsFilePath, "file", "f", "", "File path to the cloud resource results")
}
