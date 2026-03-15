package cmd

import (
	"log"
	"strings"

	"github.com/jlee3227/simple-printer/util/simple_print"
	"github.com/spf13/cobra"
)

// TODO: Add flag for reading text from a file
// textCmd represents the text command
var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Subcommand for printing text",
	Long:  `A subcommand for printing a string of text. Ideally, just a paragraph at most.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Please provide a text to print.")
		}

		text := strings.Join(args, " ")
		if err := simple_print.Print(text); err != nil {
			log.Println("Failed to print text:", err)
		}
	},
}

// TODO: Add flag for reading a list from a file
// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Subcommand for printing a bulleted list of strings",
	Long:  `A subcommand for printing a bulleted list of items. Each item is a text string. Entering an empty line will exit input and print the list.`,
	Run: func(cmd *cobra.Command, args []string) {
		var list []string

		filepath, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatal("Failed to get flag:", err)
		}

		if filepath != "" {
			list, err = simple_print.ReadListFromFile(filepath)
			if err != nil {
				log.Fatal("Failed to read list file:", err)
			}
		} else {
			list, err = simple_print.ReadListFromIO()
			if err != nil {
				log.Fatal("Failed to read list from IO:", err)
			}
		}

		if err := simple_print.PrintList(list); err != nil {
			log.Println("Failed to print list:", err)
		}
	},
}

func init() {
	listCmd.Flags().StringP("file", "f", "", "File to print list from")
	rootCmd.AddCommand(textCmd)
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// textCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// textCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
