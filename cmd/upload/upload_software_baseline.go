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
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/deathlabs/emu/config"
	"github.com/deathlabs/emu/emass"
	"github.com/deathlabs/emu/models"
	"github.com/deathlabs/emu/output"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
)

var (
	softwareBaselineFileType string
	softwareBaselinePath     string
)

var (
	uploadSoftwareBaselineCmd = &cobra.Command{
		Use:   "software-baseline",
		Short: "Upload a software baseline to eMASS",
		RunE:  uploadSoftwareBaseline,
	}
)

func uploadSoftwareBaseline(cmd *cobra.Command, args []string) error {
	var (
		body     []byte
		endpoint string
		entries  []models.SoftwareBaselineEntry
		err      error
		file     *os.File
		response *http.Response
		system   models.System
		systems  []models.System
	)

	// Filter systems based on system IDs provided via the root-level --system-ids flag.
	// If no system IDs are provided, this will return all systems for the active profile.
	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	file, err = os.Open(softwareBaselinePath)
	if err != nil {
		return err
	}
	defer file.Close()

	switch softwareBaselineFileType {
	case "json":
		err = json.NewDecoder(file).Decode(&entries)
	default:
		err = gocsv.UnmarshalFile(file, &entries)
	}
	if err != nil {
		return err
	}

	body, err = json.Marshal(entries)
	if err != nil {
		return err
	}

	// Loop through the filtered systems and upload a software baseline to each one.
	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/sw-baseline", config.Data.URL, system.ID)
		response, err = emass.Post(system.ConfigProfile, endpoint, bytes.NewBuffer(body), "application/json")
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
	// Define flags for the "emu upload software-baseline" subcommand.
	uploadSoftwareBaselineCmd.PersistentFlags().StringVarP(&softwareBaselineFileType, "type", "t", "csv", "File type of the software baseline (csv, json)")
	uploadSoftwareBaselineCmd.PersistentFlags().StringVarP(&softwareBaselinePath, "file", "f", "", "Filepath to software baseline")
}
