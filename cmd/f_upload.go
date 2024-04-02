/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fuploadCmd = &cobra.Command{
	Use:   "fupload",
	Short: "Upload files to Slack channels",
	Long: `fupload is a tool that enables you to easily upload files to your Slack channels. This command allows you to upload local files or directly from the internet. Choose the appropriate option to initiate the upload process.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📁 Melvin - Easily Upload files to your Slack channels!")
		var userInput string
		fmt.Println("What can I do for you today?")
		fmt.Printf("1. Upload Local File\n2. Upload File from URL\n")
		fmt.Scanln(&userInput)

		switch userInput {
		case "1":
			// Call local file upload function
		case "2":
			// Call fetch from URL and then upload file to Slack
		default:
			fmt.Printf("Invalid option: %s\n", userInput)
		}
	},
}


func init() {
	rootCmd.AddCommand(fuploadCmd)
	fuploadCmd.Flags().BoolP("directory", "d", false, "Include folder of directories. If not provided command will be exceuted for singular files")
	fuploadCmd.Flags().StringP("token", "t", "", "Slack API token for authentication. If not provided, the tool will check the SLACKMATE_API_TOKEN environmental variable. If neither of these values is present or if permissions are insufficient, the program will exit with an error message.")
	fuploadCmd.Flags().StringP("channel", "c", "", "Slack Workspace Channel ID. If not provided, the tool will check the SLACKMATE_DEFAULT_CHANNEL environmental variable. If neither of these values is present, the program will exit with an error message")
}
