package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simple-printer",
	Short: "Print to a receipt printer",
	Long:  `Print text, images, and lists to a receipt printer via ESC/POS.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("device", "d", "/dev/usb/lp0", "Printer device path")
}
