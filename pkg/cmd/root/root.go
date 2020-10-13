package root

import (
	"fmt"
	"os"

	pushCmd "github.com/aubinlrx/winddle-pr-cli/pkg/cmd/push"
	versionCmd "github.com/aubinlrx/winddle-pr-cli/pkg/cmd/version"
	viewCmd "github.com/aubinlrx/winddle-pr-cli/pkg/cmd/view"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "winddle-pr <command> <subcommand> [flags]",
	Short: "Winddle PR CLI",
	Long:  "Work seamlessly with Winddle Github Repository from the command line.",
}

func Execute() {
	rootCmd.AddCommand(versionCmd.NewCommandVersion())
	rootCmd.AddCommand(viewCmd.NewCommandView())
	rootCmd.AddCommand(pushCmd.NewCommandPush())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.winddle-pr.yaml)")

	rootCmd.Flags().BoolP("toogle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".winddle-pr")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
