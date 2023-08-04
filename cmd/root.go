package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/snowkluster/unseen/cmd/hash"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "unseen",
	Short: "hasing and hash detection tool",
	Long: "unseen is a hashing tool that can also detect different hashed strings",
	// hello there i am coding in python --smz coder
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "display application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 1.0.0")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(hash.HashCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.unseen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubcommandPalettes()
}


