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
	"net/url"
	"strings"

	"github.com/deathlabs/emu/emass"
	"github.com/deathlabs/emu/models"
	"github.com/deathlabs/emu/output"
	"github.com/spf13/cobra"
)

var (
	controlIDs []string
	getCmd     = &cobra.Command{
		Use:   "get",
		Short: "Get data",
	}
	getControlCmd = &cobra.Command{
		Use:   "control",
		Short: "Get data about controls",
		Run:   getControls,
	}
	getControlApprovalsCmd = &cobra.Command{
		Use:   "approvals",
		Short: "Get data about control approvals",
		Run:   getControlApprovals,
	}
	getArtifactsCmd = &cobra.Command{
		Use:   "artifacts",
		Short: "Get data about artifacts",
		Run:   getArtifacts,
	}
	getRolesCmd = &cobra.Command{
		Use:   "roles",
		Short: "Get data about system roles",
		Run:   getRoles,
	}
	getSystemsCmd = &cobra.Command{
		Use:   "systems",
		Short: "Get data about systems",
		Run:   getSystems,
	}
	getWorkflowsCmd = &cobra.Command{
		Use:   "workflows",
		Short: "Get data about workflows",
		Run:   getWorkflows,
	}
	role         string
	roleCategory string
	policy       string
)

func getArtifacts(cmd *cobra.Command, args []string) {
	var (
		endpoint string
		err      error
		response *http.Response
		system   models.System
		systems  []models.System
		profile  models.ConfigProfile
	)

	systems, err = filterSystems(config, activeProfileName, systemIDs)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, system = range systems {
		profile = system.ConfigProfile
		endpoint = fmt.Sprintf("%s/api/systems/%d/artifacts", config.URL, system.ID)

		response, err = emass.Get(profile, endpoint)
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

func getControls(cmd *cobra.Command, args []string) {
	var (
		endpoint string
		err      error
		response *http.Response
		system   models.System
		systems  []models.System
		profile  models.ConfigProfile
	)

	systems, err = filterSystems(config, activeProfileName, systemIDs)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, system = range systems {
		profile = system.ConfigProfile

		endpoint = fmt.Sprintf("%s/api/systems/%d/controls", config.URL, system.ID)

		controlAcronyms := strings.Join(controlIDs, ",")
		params := url.Values{}
		params.Set("acronyms", controlAcronyms)
		if len(params) > 0 {
			endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
		}

		response, err = emass.Get(profile, endpoint)
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

func getControlApprovals(cmd *cobra.Command, args []string) {
	var (
		endpoint string
		err      error
		response *http.Response
		system   models.System
		systems  []models.System
		profile  models.ConfigProfile
	)

	systems, err = filterSystems(config, activeProfileName, systemIDs)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, system = range systems {
		profile = system.ConfigProfile
		endpoint = fmt.Sprintf("%s/api/systems/%d/approval/cac", config.URL, system.ID)

		response, err = emass.Get(profile, endpoint)
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

func getRoles(cmd *cobra.Command, args []string) {
	var (
		endpoint string
		err      error
		params   url.Values
		profile  models.ConfigProfile
		profiles []models.ConfigProfile
		response *http.Response
	)

	profiles, err = filterProfiles(config, activeProfileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, profile = range profiles {
		endpoint = fmt.Sprintf("%s/api/system-roles", config.URL)

		if roleCategory != "" {
			if role == "" {
				fmt.Printf("profile %s: %v\n", profile.Name, "a category and role are required")
				continue
			}
			endpoint = fmt.Sprintf("%s/%s", endpoint, roleCategory)
			params = url.Values{}
			params.Set("role", role)
			if policy != "" {
				params.Set("policy", policy)
			}
			if len(params) > 0 {
				endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
			}
		}

		response, err = emass.Get(profile, endpoint)
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

func getSystems(cmd *cobra.Command, args []string) {
	var (
		endpoint string
		err      error
		profile  models.ConfigProfile
		profiles []models.ConfigProfile
		response *http.Response
		system   models.System
		systems  []models.System
	)

	// Filter by system if specific systems are requested.
	if len(systemIDs) != 0 {
		systems, err = filterSystems(config, activeProfileName, systemIDs)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, system = range systems {
			profile = system.ConfigProfile
			endpoint = fmt.Sprintf("%s/api/systems/%d", config.URL, system.ID)

			response, err = emass.Get(profile, endpoint)
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
	} else {
		// Filter by profile if no specific systems are requested.
		profiles, err = filterProfiles(config, activeProfileName)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, profile = range profiles {
			endpoint = fmt.Sprintf("%s/api/systems", config.URL)

			response, err = emass.Get(profile, endpoint)
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
}

func getWorkflows(cmd *cobra.Command, args []string) {
	var (
		err      error
		response *http.Response
		system   models.System
		systems  []models.System
		profile  models.ConfigProfile
		endpoint string
	)

	systems, err = filterSystems(config, activeProfileName, systemIDs)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, system = range systems {
		profile = system.ConfigProfile
		endpoint = fmt.Sprintf("%s/api/systems/%d/approval/pac", config.URL, system.ID)

		response, err = emass.Get(profile, endpoint)
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

func init() {
	// Define parameters for the "emu get control" command.
	getControlCmd.PersistentFlags().StringSliceVarP(&controlIDs, "control-id", "", []string{}, "Control IDs")

	// Define parameters for the "emu get roles" command.
	getRolesCmd.Flags().StringVarP(&roleCategory, "category", "", "", "PAC, CAC, or Other")
	getRolesCmd.Flags().StringVarP(&role, "role", "", "", "ISO, ISSM, SCA, Auditor, AO, etc. (required if --category is used)")
	getRolesCmd.Flags().StringVarP(&policy, "policy", "", "", "RMF, DIACAP, or Reporting")

	// Attach commands to the "emu get control" command
	getControlCmd.AddCommand(getControlApprovalsCmd)

	// Attach commands to the "emu get" command
	getCmd.AddCommand(getArtifactsCmd)
	getCmd.AddCommand(getControlCmd)
	getCmd.AddCommand(getRolesCmd)
	getCmd.AddCommand(getSystemsCmd)
	getCmd.AddCommand(getWorkflowsCmd)

	// Attach commands to the "emu" command.
	rootCmd.AddCommand(getCmd)
}
