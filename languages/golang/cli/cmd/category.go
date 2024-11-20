/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Help()
	// 	testShort, _ := cmd.Flags().GetString("testShort")
	// 	fmt.Println("Called: ", testShort)

	// 	boolT, _ := cmd.Flags().GetBool("boolT")
	// 	fmt.Println("boolT", boolT)

	// 	fmt.Println("Int", cmd.Flags().Lookup("intValue").Value)

	// 	fmt.Println("FlagTest", flagTest)
	// },
	// // PreRun: func(cmd *cobra.Command, args []string) {
	// // 	fmt.Println("PreRun")
	// // },
	// // PostRun: func(cmd *cobra.Command, args []string) {
	// // 	fmt.Println("Postrun")
	// // },
	// // RuneE: func(cmd *cobra.Command, args []string) {
	// // 	fmt.Println("RunE")
	// // }
}

//var flagTest string

func init() {
	rootCmd.AddCommand(categoryCmd)
	// Global flags *PersistentFlags* are available to all commands
	// categoryCmd.PersistentFlags().String("name", "", "Name of the course")
	// // Local flags *Flags* are only available to the command they are defined in
	// categoryCmd.Flags().String("description", "", "Description of the course")

	// //Short flag
	// //testShort is the name of the flag
	// //st is the short flag
	// //"test" is the default value[
	// //"Short Flag" is the description
	// //Short flags are defined using the StringP method
	// categoryCmd.Flags().StringP("testShort", "t", "Y", "Short Flag")

	// categoryCmd.PersistentFlags().BoolP("boolT", "b", false, "Help message for toggle")
	// categoryCmd.Flags().Int16("intValue", 0, "Integer Value")

	// // Bind the flag to a variable
	// categoryCmd.Flags().StringVarP(&flagTest, "flagTest", "f", "var", "Flag Test")

}
