package gomo_test

import (
	"fmt"
	"log"

	"github.com/moltin/gomo"
	"github.com/moltin/gomo/core"
)

func Example() {
	// create a new client with clinet credentials
	client := gomo.NewClient(
		gomo.ClientCredentials(
			"client_id",
			"client_secret",
		),
	)
	// handle an authentication error
	if err := client.Authenticate(); err != nil {
		log.Fatal(err)
	}

	// create a product
	product := core.Product{
		Name: "My new product",
	}

	// send the create request
	wrapper, err := client.Post("products", gomo.Body(product), gomo.Data(&product))
	if err != nil {
		log.Fatal(err)
	}

	// print the execution time metic
	log.Println("Execution time:", wrapper.ExecutionTime.Elapsed())

	// update a product field
	product.Name = "Updated Product"

	// send the update request
	wrapper, err = client.Put(fmt.Sprintf("products/%s", product.ID), gomo.Body(product))
	if err != nil {
		log.Fatal(err)
	}

	// delete the product
	_, err = client.Delete(fmt.Sprintf("products/%s", product.ID))
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleClientCredentials() {
	creds := gomo.ClientCredentials(
		"client_id",
		"client_secret",
	)
	gomo.NewClient(creds)
}

func ExampleImplicitCredentials() {
	creds := gomo.ImplicitCredentials(
		"client_id",
	)
	gomo.NewClient(creds)
}

func ExampleNewClient() {
	client := gomo.NewClient(
		gomo.ClientCredentials(
			"client_id",
			"client_secret",
		),
	)
	log.Println(client.APIVersion)
}
