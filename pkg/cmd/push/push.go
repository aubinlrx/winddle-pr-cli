package push

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommandPush() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "push",
		Short: "Push pull requests grouped by label to remote source (default: slack)",
		Run: func(cmd *cobra.Command, args []string) {
			pushRun()
		},
	}

	return cmd
}

func pushRun() {
	fmt.Println("Diplay push")
}
