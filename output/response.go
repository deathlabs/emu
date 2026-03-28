package output

import (
	"fmt"
	"os"
	"strings"

	"github.com/deathlabs/emu/types"
)

func Response(response types.ResponseData, format string) {
	switch strings.ToLower(format) {
	case "json":
		ToJSON(response)
	case "table":
		ToTable(response)
	case "yaml":
		ToYAML(response)
	default:
		fmt.Printf("Unsupported output format: %s\n", format)
		os.Exit(1)
	}
}
