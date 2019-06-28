package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/moltin/gomo"
	"github.com/moltin/gomo/core"
)

func main() {
	filename := "example_image.jpeg"
	clientID := os.Getenv("MOLTIN_CLIENT_ID")
	clientSecret := os.Getenv("MOLTIN_CLIENT_SECRET")

	// Instantiate a new client and provide an options function to override the default authentication method
	// Options can be found at https://github.com/moltin/gomo/blob/master/options.go
	client := gomo.NewClient(gomo.ClientCredentials(clientID, clientSecret))

	// Execute the debug option function for the client in order to turn on debugging
	client.EnableDebug()

	// Authenticate against the Moltin API
	err := client.Authenticate()
	if err != nil {
		fmt.Println(err)
	}

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		fmt.Println(err)
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println(err)
	}

	bodyWriter.Close()

	fileRequest := core.FileUploadRequest{
		Public: false,
		File:   *bodyBuf,
	}

	var fileResponse core.File

	_, err = client.Post("/files", gomo.Body(fileRequest), gomo.Data(&fileResponse))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fileResponse)
	}
}
