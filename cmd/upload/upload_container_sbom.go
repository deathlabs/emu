package upload

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
)

var (
	containerSbomContainerIdentifier string
	containerSbomContainerName       string
	containerSbomSbomPath            string
	containerSbomSbomFormat          string
)

var (
	uploadContainerSBOMCmd = &cobra.Command{
		Use:   "container-sbom",
		Short: "Upload an container sbom to eMASS",
		RunE:  uploadContainerSBOM,
	}
)

func uploadContainerSBOM(cmd *cobra.Command, args []string) error {
	var (
		body       bytes.Buffer
		endpoint   string
		err        error
		file       *os.File
		fileWriter io.Writer
		headers    map[string]string
		response   *http.Response
		system     models.System
		systems    []models.System
		writer     *multipart.Writer
	)

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		writer = multipart.NewWriter(&body)

		fileWriter, err = writer.CreateFormFile("file", filepath.Base(containerSbomSbomPath))
		if err != nil {
			return err
		}

		file, err = os.Open(containerSbomSbomPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			return err
		}

		err = file.Close()
		if err != nil {
			return err
		}

		err = writer.WriteField("containerName", containerSbomContainerName)
		if err != nil {
			return err
		}

		err = writer.WriteField("containerIdentifier", containerSbomContainerIdentifier)
		if err != nil {
			return err
		}

		err = writer.WriteField("format", containerSbomSbomFormat)
		if err != nil {
			return err
		}

		err = writer.Close()
		if err != nil {
			return err
		}

		endpoint = fmt.Sprintf("%s/api/systems/%d/containers/sbom", config.Data.URL, system.ID)

		headers = map[string]string{
			"Content-Type": writer.FormDataContentType(),
		}

		response, err = emass.Post(system.ConfigProfile, endpoint, headers, &body)
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
	uploadContainerSBOMCmd.PersistentFlags().StringVarP(&containerSbomSbomPath, "file", "f", "", "Filepath to container SBOM")
	uploadContainerSBOMCmd.PersistentFlags().StringVarP(&containerSbomSbomFormat, "format", "", "", "Container SBOM format")
	uploadContainerSBOMCmd.PersistentFlags().StringVarP(&containerSbomContainerName, "container-name", "", "", "Container name")
	uploadContainerSBOMCmd.PersistentFlags().StringVarP(&containerSbomContainerIdentifier, "container-id", "", "", "Container ID (e.g., tag)")
}
