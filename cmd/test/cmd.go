package test

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "test",
		Short: "Test a function",
	}
)

func init() {
	Cmd.AddCommand(testAPICmd)
}
