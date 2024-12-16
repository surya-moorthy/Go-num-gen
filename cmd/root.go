/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)
// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "securerandnum",
	Short: "Secure random number generator",
	Long: `An interface to generate secure random numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd)
	},
}

func Execute(){
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init(){
   rootCmd.Execute()
}