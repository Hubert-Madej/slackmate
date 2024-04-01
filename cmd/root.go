package cmd

import (
	"fmt"
	"os"

	"github.com/Hubert-Madej/slackmate/pkg/models"
	"github.com/Hubert-Madej/slackmate/pkg/utils"
	"github.com/spf13/cobra"
)

var configFileName = "config.json"

var rootCmd = &cobra.Command{
	Use:   "slackmate",
	Short: "Your personal Slack Assistant!",
	Long: `Slackmate is a CLI tool designed to streamline and automate your workflow within Slack Workspaces.
	It empowers users to effortlessly manage day-to-day tasks, such as:
			- Generating custom templates for markdown files
			- Automating repetitive actions to enhance productivity
			- Simplifying communication and collaboration among team members
	With Slackmate, you can optimize your Slack experience and focus on what matters most.
`,
	Run: func(cmd *cobra.Command, args []string) {
		encryptionKey, err := cmd.Flags().GetString("encryptionKey")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		key, err := utils.PreloadEncryptionKey(encryptionKey)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var config models.Config

		err = utils.LoadConfig(configFileName, &config, key)
		if err != nil {
			fmt.Println(err)
		}

		utils.CheckAndSetEnvVariables(&config)

		err = utils.SaveConfig(config, configFileName, key)
		if err != nil {
			fmt.Println("❌ Failed to save config file:", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("encryptionKey", "k", "", "Encryption key for configuration file (AES).")
}
