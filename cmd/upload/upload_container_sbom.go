/*
Copyright © 2026 Vic Fernandez III <@cyberphor>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package upload

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deathlabs/emu/config"
	"github.com/deathlabs/emu/emass"
	"github.com/deathlabs/emu/models"
	"github.com/deathlabs/emu/output"
	"github.com/spf13/cobra"
)

var (
	containerIdentifier string
	containerName       string
	sbomPath            string
	sbomFormat          string
)

var (
	uploadContainerSBOMCmd = &cobra.Command{
		Use:   "container-sbom",
		Short: "Upload an sbom to eMASS",
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
		response   *http.Response
		system     models.System
		systems    []models.System
		writer     *multipart.Writer
	)

	// Filter systems based on system IDs provided via the root-level --system-ids flag.
	// If no system IDs are provided, this will return all systems for the active profile.
	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	// Loop through the filtered systems and upload a container SBOM to each one.
	for _, system = range systems {
		writer = multipart.NewWriter(&body)

		fileWriter, err = writer.CreateFormFile("file", filepath.Base(sbomPath))
		if err != nil {
			return err
		}

		file, err = os.Open(sbomPath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			return err
		}

		writer.WriteField("containerName", containerName)
		writer.WriteField("containerIdentifier", containerIdentifier)
		writer.WriteField("format", sbomFormat)
		writer.Close()

		endpoint = fmt.Sprintf("%s/api/systems/%d/containers/sbom", config.Data.URL, system.ID)
		response, err = emass.Post(system.ConfigProfile, endpoint, &body, writer.FormDataContentType())
		if err != nil {
			return err
		}

		// Print the response in the specified output format.
		err = output.Response(response, config.OutputFormat)
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	// Define flags for the "emu get container-sbom" subcommand.
	uploadContainerSBOMCmd.PersistentFlags().StringVarP(&sbomPath, "file", "f", "", "Filepath to container SBOM")
	uploadContainerSBOMCmd.PersistentFlags().StringVarP(&sbomPath, "format", "", "", "Container SBOM format")
	uploadContainerSBOMCmd.PersistentFlags().StringVarP(&containerName, "container-name", "", "", "Container name")
	uploadContainerSBOMCmd.PersistentFlags().StringVarP(&containerIdentifier, "container-id", "", "", "Container ID (e.g., tag)")
}
