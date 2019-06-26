package main

import (
	"fmt"
	"os"

	"github.com/moltin/gomo"
	"github.com/moltin/gomo/core"
)

func main() {

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

	type MyProduct struct {
		core.Product
		Material         string `json:"material,omitempty"`
		MaxWatt          int    `json:"max_watt,omitempty"`
		BulbQty          int    `json:"bulb_qty,omitempty"`
		Bulb             bool   `json:"bulb,omitempty"`
		New              string `json:"new,omitempty"`
		OnSale           bool   `json:"on_sale,omitempty"`
		BackgroundColour string `json:"background_colour,omitempty"`
		Finish           string `json:"finish,omitempty"`
	}

	product := MyProduct{
		Product: core.Product{
			Type:          "product",
			Name:          "Crown",
			Slug:          "crown2",
			SKU:           "CWLP100BLK2",
			Description:   "Abstract, sculptural, refined and edgy with a modern twist. Its symmetrical, spoked structure generates a clever geometric presence, which works well in a contemporary environment.",
			ManageStock:   true,
			Status:        "live",
			CommodityType: "physical",
			Price: []core.ProductPrice{
				core.ProductPrice{
					Amount:      47500,
					Currency:    "USD",
					IncludesTax: true,
				},
			},
		},
		BackgroundColour: "#d9d9d9",
		Finish:           "test",
	}

	_, err = client.Post("/products", gomo.Body(&product), gomo.Data(&product))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(product)
	}
}
