package cmd

import (
	"strings"

	"github.com/jlee3227/simple-printer/util/simple_print"
	"github.com/spf13/cobra"
)

var textCmd = &cobra.Command{
	Use:   "text <words...>",
	Short: "Print text",
	Long:  `Print a string of text to the receipt printer.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		device, err := cmd.Flags().GetString("device")
		if err != nil {
			return err
		}
		return simple_print.Print(device, strings.Join(args, " "))
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print a bulleted list",
	Long:  `Print a bulleted list. Enter items interactively, or provide a file with -f. An empty line ends interactive input.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		device, err := cmd.Flags().GetString("device")
		if err != nil {
			return err
		}

		filepath, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}

		var list []string
		if filepath != "" {
			list, err = simple_print.ReadListFromFile(filepath)
		} else {
			list, err = simple_print.ReadListFromIO()
		}
		if err != nil {
			return err
		}

		return simple_print.PrintList(device, list)
	},
}

func init() {
	listCmd.Flags().StringP("file", "f", "", "File to read list from")
	rootCmd.AddCommand(textCmd)
	rootCmd.AddCommand(listCmd)
}
