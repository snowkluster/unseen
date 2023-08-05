/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package detect

import (
	"fmt"
	"regexp"
	"github.com/spf13/cobra"
)

// detectCmd represents the detect command
var DetectCmd = &cobra.Command{
	Use:   "detect",
	Short: "pass a string to detect the method used to hash it",
	Run: func(cmd *cobra.Command, args []string) {
		md5Pattern := "^[0-9a-fA-F]{32}$"
		match, err := regexp.MatchString(md5Pattern, args[0])
		if err != nil {
			fmt.Println("Error while matching: ",err)
			return
		}
		if match {
			fmt.Println("MD5 hash detected")
		} else {
			fmt.Println("Hash not found")
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// detectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// detectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
