package cmd

import (
	"github.com/jlee3227/simple-printer/util/simple_print"
	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:   "image <file>",
	Short: "Print an image",
	Long:  `Print an image to the receipt printer. Resizes and dithers to fit receipt width.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		device, err := cmd.Flags().GetString("device")
		if err != nil {
			return err
		}
		return simple_print.PrintImage(device, args[0])
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
}
