package cmd

import (
	"log"

	simple "github.com/jlee3227/simple-printer/util/print"
	"github.com/spf13/cobra"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image ",
	Short: "Subcommand for printing an image",
	Long:  `A subcommand for printing an image. The program will do it's best to convert the provided image into a PNG and resize so that it can fit onto receipt.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Please provide a file name.")
		}

		filename := args[0]
		if err := simple.PrintImage(filename); err != nil {
			log.Println("Failed to print:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
