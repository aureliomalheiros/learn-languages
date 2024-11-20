/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		command,_ := cmd.Flags().GetString("command")
		if command == "ping" {
			cmd.Println("ping")
		}else if command == "pong" {
			cmd.Println("pong")
		}else {
			fmt.Println("Please provide a valid")
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringP("command", "c", "", "Should ping or pong")
	testCmd.MarkFlagRequired("command")
}
