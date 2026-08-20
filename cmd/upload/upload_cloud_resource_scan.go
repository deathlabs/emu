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
	cloudResourceScanFilePath string
)

var (
	uploadCloudResourceScanCmd = &cobra.Command{
		Use:   "cloud-resource-scan",
		Short: "Upload a cloud resource scan",
		RunE:  uploadCloudResourceScan,
	}
)

func uploadCloudResourceScan(cmd *cobra.Command, args []string) error {
	var (
		cloudResourceScan     []models.CloudResourceResult
		cloudResourceScanFile []byte
		cloudResourceScanJson []byte
		err                   error
		endpoint              string
		headers               map[string]string
		request               *bytes.Buffer
		response              *http.Response
		system                models.System
		systems               []models.System
	)

	headers = map[string]string{
		"Content-Type": "application/json",
	}

	cloudResourceScanFile, err = os.ReadFile(cloudResourceScanFilePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(cloudResourceScanFile, &cloudResourceScan)
	if err != nil {
		return err
	}

	cloudResourceScanJson, err = json.Marshal(cloudResourceScan)
	if err != nil {
		return err
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/cloud-resource-results", config.Data.URL, system.ID)

		request = bytes.NewBuffer(cloudResourceScanJson)

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
	uploadCloudResourceScanCmd.PersistentFlags().StringVarP(&cloudResourceScanFilePath, "file", "f", "", "File path to the cloud resource scan")
}
