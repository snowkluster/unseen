/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package hash

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
	"github.com/spf13/cobra"
)

// sha256Cmd represents the sha256 command
var sha256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "hash a string using sha2",
	Run: func(cmd *cobra.Command, args []string) {
		hash := sha256.Sum256([]byte(args[0]))
		fmt.Println(hex.EncodeToString(hash[:]))
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sha2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sha2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
