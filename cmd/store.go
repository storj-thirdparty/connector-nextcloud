package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// storeCmd represents the store command.
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Command to upload data to storj V3 network.",
	Long:  `Command to connect to desired Nextcloud account and back-up the complete data to given Storj Bucket.`,
	Run:   nextcloudStore,
}

func init() {

	// Setup the store command with its flags.
	rootCmd.AddCommand(storeCmd)
	var defaultNextcloudFile string
	var defaultStorjFile string
	storeCmd.Flags().BoolP("accesskey", "a", false, "Connect to storj using access key(default connection method is by using API Key).")
	storeCmd.Flags().BoolP("share", "s", false, "For generating share access of the uploaded backup file.")
	storeCmd.Flags().StringVarP(&defaultNextcloudFile, "nextcloud", "n", "././config/nextcloud_property.json", "full filepath contaning Nextcloud configuration.")
	storeCmd.Flags().StringVarP(&defaultStorjFile, "storj", "u", "././config/storj_config.json", "full filepath contaning storj V3 configuration.")
}

func nextcloudStore(cmd *cobra.Command, args []string) {

	// Process arguments from the CLI.
	nextcloudConfigfilePath, _ := cmd.Flags().GetString("nextcloud")
	fullFileNameStorj, _ := cmd.Flags().GetString("storj")
	useAccessKey, _ := cmd.Flags().GetBool("accesskey")
	useAccessShare, _ := cmd.Flags().GetBool("share")

	// Read Nextcloud instance's configurations from an external file and create an Nextcloud configuration object.
	configNextcloud := LoadNextcloudProperty(nextcloudConfigfilePath)

	// Read storj network configurations from and external file and create a storj configuration object.
	storjConfig := LoadStorjConfiguration(fullFileNameStorj)

	// Connect to storj network using the specified credentials.
	access, project := ConnectToStorj(fullFileNameStorj, storjConfig, useAccessKey)

	// Connect to Nextcloud using the specified credientials
	nextcloudClient := ConnectToNextcloud(configNextcloud)

	// Create a slice of file names to be uploaded.
	GetFilesWithPaths(nextcloudClient, "/")

	fmt.Printf("\nInitiating back-up.\n")
	// Fetch all backup files from Nextcloud instance and simultaneously store them into desired Storj bucket.
	for i := 0; i < len(AllFilesWithPaths); i++ {
		file := AllFilesWithPaths[i]
		nextcloudReader := GetReader(nextcloudClient, file)
		UploadData(project, storjConfig, file, nextcloudReader, AllFilesWithPaths[i])

	}
	fmt.Printf("\nBack-up complete.\n\n")

	// Create restricted shareable serialized access if share is provided as argument.
	if useAccessShare {
		ShareAccess(access, storjConfig)
	}
}
