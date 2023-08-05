/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package detect

import (
	"fmt"
	"github.com/spf13/cobra"
)

// sha1Cmd represents the sha1 command
var sha1Cmd = &cobra.Command{
	Use:   "sha1",
	Short: "detect sha1 hashes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sha1 called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sha1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sha1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
