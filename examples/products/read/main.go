package main

import (
	"fmt"
	"log"
	"os"

	moltin "github.com/andrew-waters/moltin-go-client"
	"github.com/andrew-waters/moltin-go-client/entities"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	client, err := moltin.NewClient(
		moltin.NewImplicitCredentials(
			os.Getenv("CLIENT_ID"),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	products := []entities.Product{}
	r, err := client.Get("products", &products)

	fmt.Printf("Status Code: %d\nResource Count: %d\n", r.StatusCode, len(products))

	// uncomment this to see the actual product structs:
	// spew.Dump(products)

}
