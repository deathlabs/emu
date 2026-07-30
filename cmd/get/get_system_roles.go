package get

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
)

var (
	rolesCategory string
	rolesRole     string
	rolesPolicy   string
)

var (
	getSystemRolesCmd = &cobra.Command{
		Use:   "system-roles",
		Short: "Get data about system roles",
		RunE:  getSystemRoles,
	}
)

func getSystemRoles(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		response *http.Response
		params   url.Values
		profile  models.ConfigProfile
		profiles []models.ConfigProfile
	)

	profiles, err = config.FilterProfiles(config.Data, config.ActiveProfileName)
	if err != nil {
		return err
	}

	for _, profile = range profiles {

		endpoint = fmt.Sprintf("%s/api/system-roles", config.Data.URL)

		if rolesCategory != "" {

			if rolesRole == "" {
				return fmt.Errorf("profile %s: a category and role are required", profile.Name)
			}

			endpoint = fmt.Sprintf("%s/%s", endpoint, rolesCategory)

			params = url.Values{}
			params.Set("role", rolesRole)
			if rolesPolicy != "" {
				params.Set("policy", rolesPolicy)
			}

			if len(params) > 0 {
				endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
			}
		}

		response, err = emass.Get(profile, endpoint)
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
	getSystemRolesCmd.Flags().StringVarP(&rolesCategory, "category", "", "", "PAC, CAC, or Other")
	getSystemRolesCmd.Flags().StringVarP(&rolesRole, "role", "", "", "ISO, ISSM, SCA, Auditor, AO, etc. (required if --category is used)")
	getSystemRolesCmd.Flags().StringVarP(&rolesPolicy, "policy", "", "", "RMF, DIACAP, or Reporting")
}
