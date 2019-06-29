package gomo_test

import (
	"fmt"
	"log"

	"github.com/moltin/gomo"
	"github.com/moltin/gomo/core"
)

func Example() {
	// create a new client with client credentials
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

	var executionTime *gomo.APIExecution
	// send the create request
	err := client.Post(
		"products",
		gomo.Body(product),
		gomo.Data(&product),
		gomo.ExecutionTime(&executionTime),
	)
	if err != nil {
		log.Fatal(err)
	}

	// print the execution time metric
	log.Println("Execution time:", executionTime.Elapsed())

	// update a product field
	product.Name = "Updated Product"

	// send the update request
	err = client.Put(
		fmt.Sprintf("products/%s", product.ID),
		gomo.Body(product),
	)
	if err != nil {
		log.Fatal(err)
	}

	// delete the product
	err = client.Delete(fmt.Sprintf("products/%s", product.ID))
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
		gomo.Endpoint("http://test.example.com/"),
		gomo.Debug(),
	)
	log.Println(client.APIVersion)
}

func ExampleIterate() {
	client := gomo.NewClient()
	_ = client.Authenticate()

	gomo.Iterate(
		100,
		func(paginate gomo.RequestResource) error {
			page := []core.Product{}
			err := client.Get("product", gomo.Data(&page))
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%d products in page\n", len(page))
			return nil
		},
	)
}

func ExampleMorePages() {
	client := gomo.NewClient()
	_ = client.Authenticate()

	var orders []core.Order
	var meta core.Meta
	err := client.Get("orders", gomo.Data(&orders), gomo.Meta(&meta))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got %d orders.\n", len(orders))
	if gomo.MorePages(meta) {
		fmt.Println("There are more orders to fetch!")
	}
}

func ExampleNextPage() {
	client := gomo.NewClient()
	_ = client.Authenticate()

	// Note that Iterate() is a neater way of doing this.
	var allOrders []core.Order
	offset := 0
	limit := 100
	for {
		var page []core.Order
		var meta core.Meta
		err := client.Get(
			"orders",
			gomo.Data(&page),
			gomo.Meta(&meta),
			gomo.Paginate(offset, limit),
		)
		if err != nil {
			log.Fatal(err)
		}
		allOrders = append(allOrders, page...)
		if !gomo.MorePages(meta) {
			break
		}
		offset, limit = gomo.NextPage(meta)
	}
	fmt.Printf("Got all the orders, a total of %d\n", len(allOrders))
}

func ExampleClient_Get() {
	client := gomo.NewClient()
	_ = client.Authenticate()

	id := "96a52ef6-62c0-47ad-809d-6390d7727d49"
	type MyProduct struct {
		core.Product
		FlowField string `json:"flow_field"`
	}
	var product MyProduct
	err := client.Get(
		fmt.Sprintf("products/%s", id),
		gomo.Data(&product),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"Product %s has flow field %s\n",
		product.Product.ID,
		product.FlowField,
	)
}

func ExampleClient_Post() {
	client := gomo.NewClient()
	_ = client.Authenticate()

	type MyProduct struct {
		core.Product
		FlowField string `json:"flow_field"`
	}
	product := MyProduct{
		Product: core.Product{
			Name: "foo",
		},
		FlowField: "foo custom",
	}
	err := client.Post(
		"products",
		gomo.Body(&product),
		gomo.Data(&product),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created product %s\n", product.Product.ID)
}

func ExampleFilter() {
	client := gomo.NewClient()
	_ = client.Authenticate()

	var draftProducts []core.Product
	err := client.Get(
		"products",
		gomo.Filter("eq(status,draft)"),
		gomo.Data(&draftProducts),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Draft products:")
	for _, product := range draftProducts {
		fmt.Println(product.ID)
	}
}

func ExampleInclude() {
	client := gomo.NewClient()
	_ = client.Authenticate()

	id := "c82d2f00-bc66-4c7d-984a-8765222abb98"
	var included struct {
		Products []core.Product `json:"products"`
	}
	err := client.Get(
		"categories/"+id,
		gomo.Include("produces"),
		gomo.Included(&included),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Category %s has products:\n", id)
	for _, product := range included.Products {
		fmt.Println(product.ID)
	}
}

func ExampleSort() {
	client := gomo.NewClient()
	_ = client.Authenticate()

	var productsByName []core.Product
	err := client.Get(
		"products",
		gomo.Sort("name"),
		gomo.Data(&productsByName),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product names:")
	for _, product := range productsByName {
		fmt.Println(product.Name)
	}
}
