package cmd_test

import (
	"context"
	"fmt"
	"log"
	"os"

	"testing"

	"github.com/storj-thirdparty/connector-nextcloud/cmd"
)

func TestMongoStore(t *testing.T) {

	storjConfig := cmd.LoadStorjConfiguration("../config/storj_config_test.json")
	_, project := cmd.ConnectToStorj(storjConfig, false)

	fileReader, err := os.Open("../testFile.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Initiating back-up.\n")
	cmd.UploadData(project, storjConfig, "testFile.txt", fileReader)
	fmt.Printf("Back-up complete.\n\n")

	fmt.Printf("\nDeleting the test back-up.\n")
	ctx := context.Background()
	backups := project.ListObjects(ctx, storjConfig.Bucket, nil)
	// Loop to find the latest back-up of all the back-ups.
	for backups.Next() {
		item := backups.Item()
		_, err := project.DeleteObject(ctx, storjConfig.Bucket, item.Key)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Deleted the test back-up.\n")
}
