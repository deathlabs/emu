package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	createPoamCmd = &cobra.Command{
		Use:   "create",
		Short: "Create data",
		Run:   createPoam,
	}
)

func createPoam(cmd *cobra.Command, args []string) {
	fmt.Println("emu create poam")
}
