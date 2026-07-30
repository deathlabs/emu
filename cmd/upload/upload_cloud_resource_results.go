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
	cloudResourceResultsPath string
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
		cloudResourceResultsBytes []byte
		endpoint                  string
		err                       error
		headers                   map[string]string
		request                   *bytes.Buffer
		requestBody               []byte
		requestBodyData           models.CloudResourceResult
		response                  *http.Response
		system                    models.System
		systems                   []models.System
	)

	headers = map[string]string{
		"Content-Type": "application/json",
	}

	cloudResourceResultsBytes, err = os.ReadFile(cloudResourceResultsPath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(cloudResourceResultsBytes, &requestBodyData)
	if err != nil {
		return err
	}

	requestBody, err = json.Marshal([]models.CloudResourceResult{requestBodyData})
	if err != nil {
		return err
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/cloud-resource-results", config.Data.URL, system.ID)

		request = bytes.NewBuffer(requestBody)
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
	uploadCloudResourceResultsCmd.PersistentFlags().StringVarP(&cloudResourceResultsPath, "file", "f", "", "File path to the cloud resource results")
}
