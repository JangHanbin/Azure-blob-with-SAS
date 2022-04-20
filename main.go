package main

import (
	"azure-with-sas/SAS"
	"context"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"log"
	"net/url"
	"strings"
)

func main() {

	u := SAS.GetCredentialFromFile("config.json").BlobServiceSASURL

	ParsedURL, _ := url.Parse(u)

	credential := azblob.NewAnonymousCredential()
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	serviceURL := azblob.NewServiceURL(*ParsedURL, p)
	containerURL := serviceURL.NewContainerURL("container-sas-public-access") // Container names require lowercase
	ctx := context.Background()

	blobURL := containerURL.NewBlockBlobURL("directory/HelloWorld2.txt") // Blob names can be mixed case

	// Create New container
	_, err := containerURL.Create(ctx, azblob.Metadata{}, azblob.PublicAccessBlob)
	if err != nil {
		println("CREATE ERROR")
		log.Fatal(err)

	}

	// Create the blob with string (plain text) content.
	data := "Hello World! play with me!"
	_, err = blobURL.Upload(ctx, strings.NewReader(data), azblob.BlobHTTPHeaders{ContentType: "text/plain"}, azblob.Metadata{}, azblob.BlobAccessConditions{}, azblob.DefaultAccessTier, nil, azblob.ClientProvidedKeyOptions{})
	if err != nil {
		log.Fatal(err)
	}

}
