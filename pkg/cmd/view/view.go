package view

import (
	"fmt"
	"os"

	"github.com/aubinlrx/winddle-pr-cli/api"
	"github.com/spf13/cobra"
)

func NewCommandView() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view",
		Short: "Display pull requests grouped by label",
		Run: func(cmd *cobra.Command, args []string) {
			viewRun()
		},
	}

	return cmd
}

func viewRun() {
	bodyString, err := api.GetAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(bodyString)
}
