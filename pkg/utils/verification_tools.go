package utils

import (
	"fmt"
	"os"

	"github.com/Hubert-Madej/slackmate/pkg/constants"
	"github.com/Hubert-Madej/slackmate/pkg/models"
)

func CheckAndSetEnvVariables(config *models.Config) {
	apiToken := os.Getenv(constants.SLACKMATE_API_TOKEN)
	channel := os.Getenv(constants.SLACKMATE_DEFAULT_CHANNEL)

	if (apiToken == "" && config.APIToken == "") || (channel == "" && config.DefaultChannel == "") {
		fmt.Println("To ensure smooth operation of the tool, please provide the following information:")
	}

	if apiToken == "" {
		if config.APIToken == "" {
			fmt.Print("❓ Please enter your Slackmate API token (SLACKMATE_API_TOKEN): ")
			fmt.Scanln(&apiToken)
		} else {
			apiToken = config.APIToken
		}
		os.Setenv(constants.SLACKMATE_API_TOKEN, apiToken)
	}

	if channel == "" {
		if config.DefaultChannel == "" {
			fmt.Print("❓ Please enter your Slackmate default channel (SLACKMATE_DEFAULT_CHANNEL): ")
			fmt.Scanln(&channel)
		} else {
			channel = config.DefaultChannel
		}
		os.Setenv(constants.SLACKMATE_DEFAULT_CHANNEL, channel)
	}

	config.APIToken = apiToken
	config.DefaultChannel = channel
}
