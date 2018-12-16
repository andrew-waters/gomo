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
	var err error

	client, err := gomo.NewClient(
		gomo.NewClientCredentials(
			os.Getenv("CLIENT_ID"),
			os.Getenv("CLIENT_SECRET"),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	// create a category
	category := makeCategoryStruct()
	wrapper, err := client.Post("categories", &category)

	// create a product
	product := makeProductStruct()
	wrapper, err = client.Post("products", &product)

	// create a relationship
	rels := make([]entities.Relationship, 0)
	rels = append(rels, entities.Relationship{
		Type: "category",
		ID:   category.ID,
	})
	wrapper, err = client.Post(fmt.Sprintf("products/%s/relationships/categories", product.ID), rels)

	// get the product and include the categories
	included := entities.ProductIncludes{}
	wrapper, err = client.Get(fmt.Sprintf("products/%s?include=categories", "f985b65a-8f20-44ed-b67f-95d876840eaa"), &product, &included)
	fmt.Printf("Status Code: %d\n", wrapper.StatusCode)
	fmt.Printf("Time Taken: %v\n", wrapper.ExecutionTime.Elapsed())
	spew.Dump(product)
	spew.Dump(included)

}

func makeProductStruct() entities.Product {

	// make the product price
	price := make([]entities.ProductPrice, 0)
	price = append(price, entities.ProductPrice{
		Amount:      10000,
		Currency:    "USD",
		IncludesTax: true,
	})

	u := unique()

	// create the product struct
	product := entities.Product{
		Name:          "A new product",
		Slug:          fmt.Sprintf("product-slug-%d", u),
		SKU:           fmt.Sprintf("sku.product.%d", u),
		Description:   "A product description",
		ManageStock:   false,
		Price:         price,
		Status:        "live",
		CommodityType: "physical",
	}

	return product
}

func unique() int64 {
	now := time.Now()
	return now.UnixNano()
}

func makeCategoryStruct() entities.Category {

	// create the category struct
	category := entities.Category{
		Name:        "A new category",
		Slug:        fmt.Sprintf("category-slug-%d", unique()),
		Description: "A category description",
		Status:      "live",
	}

	return category
}
