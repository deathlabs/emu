package update

import (
	"bytes"
	"encoding/json"
	"errors"
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
	artifactFilePath string
)

var (
	updateArtifactCmd = &cobra.Command{
		Use:   "artifact",
		Short: "Update artifact data",
		RunE:  updateArtifact,
	}
)

func updateArtifact(cmd *cobra.Command, args []string) error {
	var (
		artifact     models.ArtifactUpdate
		artifactData []byte
		artifacts    []models.ArtifactUpdate
		endpoint     string
		err          error
		headers      map[string]string
		request      *bytes.Buffer
		response     *http.Response
		system       models.System
		systems      []models.System
	)

	if artifactFilePath == "" {
		return errors.New("an argument for the 'file' parameter is required")
	}

	artifactData, err = os.ReadFile(artifactFilePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(artifactData, &artifacts)
	if err != nil {
		return err
	}

	if len(artifacts) == 0 {
		return errors.New("the artifact update file must contain at least one artifact")
	}

	for _, artifact = range artifacts {
		if artifact.Filename == "" {
			return errors.New("each artifact update must include a filename")
		}
		if artifact.Type == "" {
			return errors.New("each artifact update must include a type")
		}
		if artifact.Category == "" {
			return errors.New("each artifact update must include a category")
		}
	}

	artifactData, err = json.Marshal(artifacts)
	if err != nil {
		return err
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	headers = map[string]string{
		"Content-Type": "application/json",
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/artifacts", config.Data.URL, system.ID)
		request = bytes.NewBuffer(artifactData)

		response, err = emass.Put(system.ConfigProfile, endpoint, headers, request)
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
	updateArtifactCmd.PersistentFlags().StringVarP(&artifactFilePath, "file", "f", "", "File path to the artifact update data")
}
