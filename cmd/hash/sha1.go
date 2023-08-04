/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package hash

import (
	"fmt"
	"crypto/sha1"
	"encoding/hex"
	"github.com/spf13/cobra"
)

// sha1Cmd represents the sha1 command
var sha1Cmd = &cobra.Command{
	Use:   "sha1",
	Short: "hash a string using md5",
	Run: func(cmd *cobra.Command, args []string) {
		hash := sha1.Sum([]byte(args[0]))
		fmt.Println(hex.EncodeToString(hash[:]))
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
