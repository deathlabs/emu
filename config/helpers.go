package config

import (
	"fmt"

	"github.com/deathlabs/emu/models"
)

func ContainsSystemID(ids []int, id int) bool {
	var current int

	// Loop through the list of system IDs and check if the provided ID is in the list.
	for _, current = range ids {
		if current == id {
			return true
		}
	}
	return false
}

func FilterProfiles(config models.Config, activeProfileName string) ([]models.ConfigProfile, error) {
	var (
		profile  models.ConfigProfile
		profiles []models.ConfigProfile
	)

	// If no active profile name is provided, return all profiles.
	if activeProfileName == "" {
		return config.ConfigProfiles, nil
	}

	// Loop through profiles in the config and filter based on the active profile name.
	for _, profile = range config.ConfigProfiles {
		if profile.Name == activeProfileName {
			profiles = append(profiles, profile)
			return profiles, nil
		}
	}

	return nil, fmt.Errorf("no profile found for name %s", activeProfileName)
}

func FilterSystems(config models.Config, profileName string, systemIDs []int) ([]models.System, error) {
	var (
		filteredSystems []models.System
		system          models.System
	)

	// Loop through systems in the config and filter based on profile name and system IDs.
	for _, system = range config.Systems {
		// If a profile name is provided and the system's profile does not match, skip this system.
		if profileName != "" && system.ConfigProfile.Name != profileName {
			continue
		}

		// If system IDs are provided and the system's ID is not in the list, skip this system.
		if len(systemIDs) > 0 && !ContainsSystemID(systemIDs, system.ID) {
			continue
		}

		// Add the system to the list of filtered systems.
		filteredSystems = append(filteredSystems, system)
	}

	// If no systems matched the filters, return an error.
	if len(filteredSystems) == 0 {
		return nil, fmt.Errorf("no systems matched the requested filters")
	}

	// Return the systems filtered based on the provided criteria.
	return filteredSystems, nil
}
