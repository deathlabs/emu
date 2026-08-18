package config

import (
	"github.com/deathlabs/emu/v4/models"
)

var (
	ActiveProfileName     string
	Data                  models.Config
	DefaultConfigFilePath string = ".emu.yaml"
	DefaultOutputFormat   string = "json"
	Filename              string
	OutputFormat          string
	SystemIDs             []int
)
