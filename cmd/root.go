package cmd

import (
	"fmt"
	"os"

	"github.com/Hubert-Madej/slackmate/pkg/constants"
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
		key, err := utils.PreloadEncryptionKey(selectKeyForExecution(cmd))
		fatalOnError(err)

		var config models.Config

		err = utils.LoadConfig(configFileName, &config, key)
		if err != nil {
			fmt.Println(err)
		}

		utils.CheckAndSetEnvVariables(&config)

		err = utils.SaveConfig(config, configFileName, key)
		fatalOnError(err)
	},
}

func Execute() {
	err := rootCmd.Execute()
	fatalOnError(err)
}

func init() {
	rootCmd.PersistentFlags().StringP("encryptionKey", "k", "", "Encryption key for configuration file (AES).")
}

func selectKeyForExecution(cmd *cobra.Command) string {
	encryptionKey, err := cmd.Flags().GetString("encryptionKey")
	fatalOnError(err)

	if encryptionKey != "" {
		return encryptionKey
	}

	return os.Getenv(constants.SLACKMATE_ENCRYPTION_KEY)
}

func fatalOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
