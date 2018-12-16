package main

import (
	"fmt"
	"log"
	"os"

	moltin "github.com/andrew-waters/moltin-go-client"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	// Create a new client with Client Credentials
	client, err := moltin.NewClient(
		moltin.NewClientCredentials(
			os.Getenv("CLIENT_ID"),
			os.Getenv("CLIENT_SECRET"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	out(client)

	// Create a new client with Implicit Credentials
	client, err = moltin.NewClient(
		moltin.NewImplicitCredentials(
			os.Getenv("CLIENT_ID"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	out(client)

}

func out(client moltin.Client) {
	fmt.Printf("`%s` client created:\nAccess token - `%s`\n\n", client.GrantType(), client.AccessToken)
}
