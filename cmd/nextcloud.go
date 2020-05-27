package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud"
)

// ConfigNextcloud defines the variables and types for login.
type ConfigNextcloud struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoadNextcloudProperty reads and parses the JSON file
// that contains a Nextcloud instance's properties.
// It returns all the properties as an object.
func LoadNextcloudProperty(fullFileName string) ConfigNextcloud {

	var configNextcloud ConfigNextcloud

	// Open and read the file.
	fileHandle, err := os.Open(filepath.Clean(fullFileName))
	if err != nil {
		log.Fatal(err)
	}

	jsonParser := json.NewDecoder(fileHandle)
	if err = jsonParser.Decode(&configNextcloud); err != nil {
		log.Fatal(err)
	}

	// Display Information about Nextcloud Instance.
	fmt.Println("Read Nextcloud configuration from the ", fullFileName, " file")
	fmt.Println("URL\t", configNextcloud.URL)
	fmt.Println("Username \t", configNextcloud.Username)
	fmt.Println("Password \t", configNextcloud.Password)

	return configNextcloud
}

//ConnectToNextcloud connects to Nextcloud by creating a new authenticated Nextcloud client.
// It returns a Nextcloud client instance to perform file transfer(back-up).
func ConnectToNextcloud(configNextcloud ConfigNextcloud) *gonextcloud.Client {

	fmt.Println("Connecting to Nextcloud...")

	// Create nextcloud client.
	nextcloudClient, err := gonextcloud.NewClient(configNextcloud.URL)
	if err != nil {
		log.Fatal("Client creation error:", err)
	}

	// Login to the nextcloud account using the specified credentials.
	if err = nextcloudClient.Login(configNextcloud.Username, configNextcloud.Password); err != nil {
		log.Fatal("Login Error", err)
	}

	return nextcloudClient
}

// AllFilesWithPaths is a list to store complete path of all the files in the Nextcloud.
// It will be used for direct transfer of files from Nextcloud to Storj.
var AllFilesWithPaths []string

// GetFilesWithPaths retrieve the files' names with the exact
// file structure from Nextcloud Server to the System
func GetFilesWithPaths(nextcloudClient *gonextcloud.Client, path string) {

	if path[len(path)-1] == '/' {
		folders, err := nextcloudClient.WebDav().ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range folders {
			if file.IsDir() {
				GetFilesWithPaths(nextcloudClient, path+file.Name()+"/")
			} else {
				AllFilesWithPaths = append(AllFilesWithPaths, path+file.Name())
			}
		}
	} else {
		AllFilesWithPaths = append(AllFilesWithPaths, path)
	}
}

// GetReader returns a Reader of corresponding file whose path is specified.
// io.ReadCloser type of object returned is used to perform transfer of file to Storj.
func GetReader(nextcloudClient *gonextcloud.Client, path string) io.ReadCloser {

	nextcloudReader, err := nextcloudClient.WebDav().ReadStream(path)
	if err != nil {
		log.Fatal("Read file error: ", err)
	}

	return nextcloudReader
}
