package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommandVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("1.0.0")
		},
	}

	return cmd
}
