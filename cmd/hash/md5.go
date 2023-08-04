/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package hash

import (
	"fmt"
	"github.com/spf13/cobra"
	"crypto/md5"
	"encoding/hex"
)

// md5Cmd represents the md5 command
var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "hash a string using md5",
	Run: func(cmd *cobra.Command, args []string) {
		hash := md5.Sum([]byte(args[0]))
		fmt.Println(hex.EncodeToString(hash[:]))
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// md5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// md5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
