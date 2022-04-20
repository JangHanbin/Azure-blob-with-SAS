package SAS

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	ConnectionString   string
	SASToken           string
	BlobServiceSASURL  string
	FileServiceSASURL  string
	QueueServiceSASURL string
	TableServiceSASURL string
}

func GetCredentialFromFile(path string) Configuration { //may function name is too long
	// From the file, get your Storage account's name and account key.
	file, fileErr := os.Open(path)
	if fileErr != nil {
		log.Fatalf("Config file open failure: %+v", fileErr)
		panic(fileErr)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	configuration := Configuration{}

	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatalf("Config parsing failure: %+v", err)
		panic(err)
	}

	return configuration

}
