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
package cmd

import (
	"fmt"
	"net/http"

	"github.com/deathlabs/emu/emass"
	"github.com/deathlabs/emu/models"
	"github.com/deathlabs/emu/output"
	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get evidence from eMASS",
	}
)

func getArtifacts(cmd *cobra.Command, args []string) {
	var (
		err      error
		response *http.Response
		system   models.System
		systems  []models.System
		profile  models.ConfigProfile
		url      string
	)

	systems, err = filterSystems(config, activeProfileName, systemIDs)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, system = range systems {
		profile = system.ConfigProfile
		url = fmt.Sprintf("%s/api/systems/%d/artifacts", config.URL, system.ID)

		response, err = emass.Get(profile, url)
		if err != nil {
			fmt.Printf("system %d: %v\n", system.ID, err)
			continue
		}

		err = output.Response(response, outputFormat)
		response.Body.Close()
		if err != nil {
			fmt.Printf("system %d: %v\n", system.ID, err)
			continue
		}
	}
}

func getSystems(cmd *cobra.Command, args []string) {
	var (
		err      error
		profile  models.ConfigProfile
		profiles []models.ConfigProfile
		response *http.Response
		url      string
	)

	profiles, err = filterProfiles(config, activeProfileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, profile = range profiles {
		url = fmt.Sprintf("%s/api/systems", config.URL)

		response, err = emass.Get(profile, url)
		if err != nil {
			fmt.Printf("profile %s: %v\n", profile.Name, err)
			continue
		}

		err = output.Response(response, outputFormat)
		response.Body.Close()
		if err != nil {
			fmt.Printf("profile %s: %v\n", profile.Name, err)
			continue
		}
	}
}

func init() {
	getCmd.AddCommand(&cobra.Command{
		Use:   "artifacts",
		Short: "Get data about one or more artifacts",
		Run:   getArtifacts,
	})
	getCmd.AddCommand(&cobra.Command{
		Use:   "systems",
		Short: "Get data about one or more systems",
		Run:   getSystems,
	})
	rootCmd.AddCommand(getCmd)
}
