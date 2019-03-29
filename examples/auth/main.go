package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andrew-waters/gomo"
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
	client := gomo.NewClient(
		gomo.NewClientCredentials(
			os.Getenv("CLIENT_ID"),
			os.Getenv("CLIENT_SECRET"),
		),
	)
	if err := client.Authenticate(); err != nil {
		log.Fatal(err)
	}
	out(client)

	// Create a new client with Implicit Credentials
	client = gomo.NewClient(
		gomo.NewImplicitCredentials(
			os.Getenv("CLIENT_ID"),
		),
	)
	if err := client.Authenticate(); err != nil {
		log.Fatal(err)
	}
	out(client)

}

func out(client gomo.Client) {
	fmt.Printf("`%s` client created:\nAccess token - `%s`\n\n", client.GrantType(), client.AccessToken)
}
