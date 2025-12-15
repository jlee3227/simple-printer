/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strings"

	simple "github.com/jlee3227/simple-printer/util/print"
	"github.com/spf13/cobra"
)

// textCmd represents the text command
var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Subcommand for printing text",
	Long:  `A subcommand for printing text strings.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Please provide a text to print.")
		}

		text := strings.Join(args, " ")
		if err := simple.Print(text); err != nil {
			log.Println("Failed to print:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(textCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// textCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// textCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
