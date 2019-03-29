package gomo_test

import (
	"fmt"
	"log"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
)

func Example() {
	// create a new client with clinet credentials
	client := gomo.NewClient(
		gomo.NewClientCredentials(
			"client_id",
			"client_secret",
		),
	)
	// handle an authentication error
	if err := client.Authenticate(); err != nil {
		log.Fatal(err)
	}

	// create a product
	product := entities.Product{
		Name: "My new product",
	}

	// send the create request
	wrapper, err := client.Post("products", &product)
	if err != nil {
		log.Fatal(err)
	}

	// print the execution time metic
	log.Println("Execution time:", wrapper.ExecutionTime.Elapsed())

	// update a product field
	product.Name = "Updated Product"

	// send the update request
	wrapper, err = client.Put(fmt.Sprintf("products/%s", product.ID), &product)
	if err != nil {
		log.Fatal(err)
	}

	// delete the product
	_, err = client.Delete(fmt.Sprintf("products/%s", product.ID))
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleNewClientCredentials() {
	creds := gomo.NewClientCredentials(
		"client_id",
		"client_secret",
	)
	gomo.NewClient(creds)
}

func ExampleNewImplicitCredentials() {
	creds := gomo.NewImplicitCredentials(
		"client_id",
	)
	gomo.NewClient(creds)
}

func ExampleNewClient() {
	client := gomo.NewClient(
		gomo.NewClientCredentials(
			"client_id",
			"client_secret",
		),
	)
	log.Println(client.APIVersion)
}

func ExampleNewClientWithCustomEndpoint() {
	client := gomo.NewClientWithCustomEndpoint(
		gomo.NewClientCredentials(
			"client_id",
			"client_secret",
		),
		"https://alt.domain.tld",
	)
	log.Println(client.APIVersion)
}
