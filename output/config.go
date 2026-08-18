package output

import (
	"fmt"
	"strings"

	"github.com/deathlabs/emu/v4/models"
)

func Config(config models.Config, format string) error {
	switch strings.ToLower(format) {
	case "json":
		ToJSON(config)
	case "yaml":
		ToYAML(config)
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}
	return nil
}
