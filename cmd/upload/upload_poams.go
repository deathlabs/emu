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
	poamsFilePath string
)

var (
	uploadPoamsCmd = &cobra.Command{
		Use:   "poams",
		Short: "Upload Plan of Action and Milestones (POA&M) items",
		RunE:  uploadPoams,
	}
)

func uploadPoams(cmd *cobra.Command, args []string) error {
	var (
		poams     []models.POAM
		poamsFile []byte
		poamsJson []byte
		err       error
		endpoint  string
		headers   map[string]string
		request   *bytes.Buffer
		response  *http.Response
		system    models.System
		systems   []models.System
	)

	headers = map[string]string{
		"Content-Type": "application/json",
	}

	poamsFile, err = os.ReadFile(poamsFilePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(poamsFile, &poams)
	if err != nil {
		return err
	}

	poamsJson, err = json.Marshal(poams)
	if err != nil {
		return err
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/poams", config.Data.URL, system.ID)

		request = bytes.NewBuffer(poamsJson)

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
	uploadPoamsCmd.PersistentFlags().StringVarP(&poamsFilePath, "file", "f", "", "File path to the POA&M")
}
