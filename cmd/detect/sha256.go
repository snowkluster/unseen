/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package detect

import (
	"fmt"
	"strings"
	"crypto/sha256"
	"encoding/hex"
	"github.com/spf13/cobra"
)

// sha256Cmd represents the sha256 command
var sha256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "detect a sha256 hash",
	Run: func(cmd *cobra.Command, args []string) {
		input := strings.TrimSpace(args[0])
		decoded, err := hex.DecodeString(input)
		if err == nil && len(decoded) == sha256.Size {
			fmt.Println("Valid SHA-256 hash!")
		} else {
			fmt.Println("Not a valid SHA-256 hash.")
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sha256Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sha256Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
