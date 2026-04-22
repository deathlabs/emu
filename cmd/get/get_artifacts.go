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
package get

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
)

var (
	artifactsFilename             string
	artifactsControlAcronyms      []string
	artifactsAssessmentProcedures []string
	artifactsCcis                 []string
	artifactsSystemOnly           bool
)

var (
	getArtifactsCmd = &cobra.Command{
		Use:   "artifacts",
		Short: "Get data about artifacts",
		RunE:  getArtifacts,
	}
)

func getArtifacts(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		params   url.Values
		response *http.Response
		system   models.System
		systems  []models.System
	)

	params = url.Values{}

	if len(artifactsAssessmentProcedures) > 0 {
		params.Set("assessmentProcedures", strings.Join(artifactsAssessmentProcedures, ","))
	}

	if len(artifactsCcis) > 0 {
		params.Set("ccis", strings.Join(artifactsCcis, ","))
	}

	if len(artifactsControlAcronyms) > 0 {
		params.Set("controlAcronyms", strings.Join(artifactsControlAcronyms, ","))
	}

	if artifactsFilename != "" {
		params.Set("filename", artifactsFilename)
	}

	if artifactsSystemOnly {
		params.Set("systemOnly", "true")
	}

	// Filter systems based on system IDs provided via the root-level --system-ids flag.
	// If no system IDs are provided, this will return all systems for the active profile.
	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	// Loop through the filtered systems and get artifacts data for each one.
	for _, system = range systems {
		// Define the endpoint for getting artifacts data for the current system.
		endpoint = fmt.Sprintf("%s/api/systems/%d/artifacts", config.Data.URL, system.ID)

		if len(params) > 0 {
			endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
		}

		response, err = emass.Get(system.ConfigProfile, endpoint)
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
	// Define flags for the "emu get artifacts" subcommand.
	getArtifactsCmd.PersistentFlags().StringVarP(&artifactsFilename, "filename", "f", "", "Filename")
	getArtifactsCmd.PersistentFlags().StringSliceVarP(&artifactsControlAcronyms, "control-acronyms", "", []string{}, "Control acronyms")
	getArtifactsCmd.PersistentFlags().StringSliceVarP(&artifactsAssessmentProcedures, "assessment-procedures", "", []string{}, "Assessment procedures")
	getArtifactsCmd.PersistentFlags().StringSliceVarP(&artifactsCcis, "ccis", "", []string{}, "CCIs")
	getArtifactsCmd.PersistentFlags().BoolVarP(&artifactsSystemOnly, "system-only", "", false, "Exclude control and AP-level artifacts only")
}
