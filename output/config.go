package output

import (
	"fmt"
	"os"
	"strings"

	"github.com/deathlabs/emu/models"
)

func Config(config models.Config, format string) {
	switch strings.ToLower(format) {
	case "json":
		ToJSON(config)
	case "table":
		ToTable(config)
	case "yaml":
		ToYAML(config)
	default:
		fmt.Printf("Unsupported output format: %s\n", format)
		os.Exit(1)
	}
}
