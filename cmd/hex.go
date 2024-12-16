/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type MyStringBuilder struct {
	strings.Builder
}

// hexCmd represents the hex command
var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Generates hex numbers",
	Long:  `An . method to give hex number by given length`,
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		hex, err := generateHexString(length)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(hex)
		}
	},
}

func generateHexString(length int) (string, error) {
	// Calculate the number of bytes needed
	byteLength := (length + 1) / 2

	// Create a byte slice to hold random data
	bytes := make([]byte, byteLength)

	// Fill the byte slice with random data
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Convert the random bytes to a hexadecimal string and trim to the desired length
	hexString := hex.EncodeToString(bytes)[:length]

	return hexString, nil
}

func init() {
	rootCmd.AddCommand(hexCmd)
	hexCmd.Flags().IntP("length", "l", 4, "length of hex")
	// hexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
