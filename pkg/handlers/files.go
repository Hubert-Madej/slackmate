package handlers

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UploadLocalFile(cmd *cobra.Command, args []string) error {
	dFlag, err := cmd.Flags().GetBool("d")
	if err != nil {
		return err
	}

	fileArr := []string{}

	// If true, then arguments provided to command are directories that contains file to upload. In other case each argument is unique file to upload.
	if dFlag {
		for _, dir := range args {
			entries, err := os.ReadDir(dir)
			if err != nil {
				return err
			}

			for _, e := range entries {
				fileArr = append(fileArr, fmt.Sprintf("%s/%s", dir, e.Name()))
			}
		}
	} else {
		for _, file := range args {
			if _, err := os.Stat(file); err != nil {
				continue
			}

			fileArr = append(fileArr, file)
		}
	}
	
	return nil
}

func UploadFileFromWeb(cmd *cobra.Command, args []string) error {
	return nil
}
