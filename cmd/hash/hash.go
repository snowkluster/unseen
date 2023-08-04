/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package hash

import (
	"github.com/spf13/cobra"
)

// hashCmd represents the hash command
var HashCmd = &cobra.Command{
	Use:   "hash",
	Short: "hash strings by passing subcommands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hashCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
