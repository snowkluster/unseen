/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package detect

import (
	"fmt"
	"strings"
	"encoding/hex"
	"github.com/spf13/cobra"
)

// sha1Cmd represents the sha1 command
var sha1Cmd = &cobra.Command{
	Use:   "sha1",
	Short: "detect sha1 hashes",
	Run: func(cmd *cobra.Command, args []string) {
		input := strings.TrimSpace(args[0])
		decoded, err := hex.DecodeString(input)
		if err != nil || len(decoded) != 20 {
			fmt.Println("Not a valid SHA-1 hash.")
		} else {
			fmt.Println("Valid SHA-1 hash!")
		}
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
