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
	staticCodeScanFilePath string
)

var (
	uploadStaticCodeScanCmd = &cobra.Command{
		Use:   "static-code-scan",
		Short: "Upload a static code scan",
		RunE:  uploadStaticCodeScan,
	}
)

func uploadStaticCodeScan(cmd *cobra.Command, args []string) error {
	var (
		staticCodeScan     []models.StaticCodeScan
		staticCodeScanFile []byte
		staticCodeScanJson []byte
		err                error
		endpoint           string
		headers            map[string]string
		request            *bytes.Buffer
		response           *http.Response
		system             models.System
		systems            []models.System
	)

	headers = map[string]string{
		"Content-Type": "application/json",
	}

	staticCodeScanFile, err = os.ReadFile(staticCodeScanFilePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(staticCodeScanFile, &staticCodeScan)
	if err != nil {
		return err
	}

	staticCodeScanJson, err = json.Marshal(staticCodeScan)
	if err != nil {
		return err
	}

	fmt.Println(string(staticCodeScanJson))

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/static-code-scans", config.Data.URL, system.ID)

		request = bytes.NewBuffer(staticCodeScanJson)

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
	uploadStaticCodeScanCmd.PersistentFlags().StringVarP(&staticCodeScanFilePath, "file", "f", "", "File path to the static code scan")
}
