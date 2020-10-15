package view

import (
	"fmt"
	"io"
	"os"

	"github.com/aubinlrx/winddle-pr-cli/api"
	"github.com/aubinlrx/winddle-pr-cli/utils"
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
	out := io.Out
	pullRequests, err := api.GetAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, pullRequest := range *pullRequests {
		fmt.Fprintf(out, "#%s", utils.Green(string(pullRequest.Number)))
		fmt.Sprintf(pullRequest.Title)
	}
}
