package upload

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
)

var (
	artifactPath   string
	artifactIsBulk bool
)

var (
	uploadArtifactCmd = &cobra.Command{
		Use:   "artifact",
		Short: "Upload an artifact to eMASS",
		RunE:  uploadArtifact,
	}
)

func uploadArtifact(cmd *cobra.Command, args []string) error {
	var (
		body            bytes.Buffer
		endpoint        string
		err             error
		file            *os.File
		fileWriter      io.Writer
		headers         map[string]string
		queryParameters url.Values
		response        *http.Response
		system          models.System
		systems         []models.System
		writer          *multipart.Writer
	)

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	if artifactPath == "" {
		return errors.New("an argument for the 'file' parameter is required")
	}

	for _, system = range systems {
		writer = multipart.NewWriter(&body)

		fileWriter, err = writer.CreateFormFile("file", filepath.Base(artifactPath))
		if err != nil {
			return err
		}

		file, err = os.Open(artifactPath)
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

		err = writer.Close()
		if err != nil {
			return err
		}

		queryParameters = url.Values{}
		queryParameters.Set("isBulk", strconv.FormatBool(artifactIsBulk))
		endpoint = fmt.Sprintf("%s/api/systems/%d/artifacts?%s", config.Data.URL, system.ID, queryParameters.Encode())

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
	uploadArtifactCmd.PersistentFlags().StringVarP(&artifactPath, "file", "f", "", "Filepath to artifact")
	uploadArtifactCmd.PersistentFlags().BoolVarP(&artifactIsBulk, "is-bulk", "", false, "When set to true, a .zip file is expected which can contain multiple artifacts")
}
