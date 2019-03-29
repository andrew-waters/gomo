package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	client := gomo.NewClient(
		gomo.NewClientCredentials(
			os.Getenv("CLIENT_ID"),
			os.Getenv("CLIENT_SECRET"),
		),
	)

	if err := client.Authenticate(); err != nil {
		log.Fatal(err)
	}

	price := make([]entities.ProductPrice, 0)
	price = append(price, entities.ProductPrice{
		Amount:      10000,
		Currency:    "USD",
		IncludesTax: true,
	})

	now := time.Now()
	unique := now.UnixNano()

	product := entities.Product{
		Name:          "A new product",
		Slug:          fmt.Sprintf("product-slug-%d", unique),
		SKU:           fmt.Sprintf("sku.product.%d", unique),
		Description:   "A product description",
		ManageStock:   false,
		Price:         price,
		Status:        "live",
		CommodityType: "physical",
	}

	r, _ := client.Post("products", &product)
	r, _ = client.Get(fmt.Sprintf("products/%s", product.ID), &product)
	fmt.Printf("Status Code: %d\n", r.StatusCode)
	fmt.Printf("Time Taken: %v\n", r.ExecutionTime.Elapsed())
	spew.Dump(product)
	spew.Dump(client.Logs)

}
