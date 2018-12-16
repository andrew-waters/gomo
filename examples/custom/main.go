package main

import (
	"fmt"
	"log"
	"os"

	moltin "github.com/andrew-waters/moltin-go-client"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// YourProductStruct is a struct to unmarshal the API response onto
// This allows you to add your custom flow fields to a struct which then replaces
// the internal struct with your own to get your data out safely.
type YourProductStruct struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	CustomAttribute string `json:"custom_attribute"`
}

func main() {

	// create an implicit client
	client, err := moltin.NewClient(
		moltin.NewImplicitCredentials(
			os.Getenv("CLIENT_ID"),
		),
	)

	// handle client error
	if err != nil {
		log.Fatal(err)
	}

	products := []YourProductStruct{}
	wrapper, err := client.Get("products", &products)

	// handle client request error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status Code: %d\nResource Count: %d\n", wrapper.StatusCode, len(products))

	// output the products (notice that they are []YourProductStruct{})
	spew.Dump(products)

}
