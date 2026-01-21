package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jlee3227/simple-printer/util/simple_print"
	"github.com/spf13/cobra"
)

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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Subcommand for printing a bulleted list of strings",
	Long:  `A subcommand for printing a bulleted list of items. Each item is a text string. Entering an empty line will exit input and print the list.`,
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string

		fmt.Println("Enter list title:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			log.Fatal("Failed to get input:", err)
		}
		lines = append(lines, scanner.Text())

		fmt.Println("Input list items. Inputting an empty line will print the list.")
		for {
			scanner.Scan()
			line := scanner.Text()
			if len(line) == 0 {
				break
			}
			lines = append(lines, line)
		}

		err := scanner.Err()
		if err != nil {
			log.Fatal("Failed to get input:", err)
		}

		if err := simple_print.PrintList(lines); err != nil {
			log.Println("Failed to print list:", err)
		}
	},
}

func init() {
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
