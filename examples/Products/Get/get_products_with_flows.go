package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/moltin/gomo"
	"github.com/moltin/gomo/core"
)

func main() {

	// Struct which includes flow fields on products and embeds the standard core fields
	type productFlowData struct {
		core.Product
		NavigationClass  string `json:"navigation_class,omitempty"`
		SpecialPriceFrom string `json:"special_price_from_date,omitempty"`
	}

	clientID := os.Getenv("MOLTIN_CLIENT_ID")

	// Instantiate a new client and provide an options function to override the default authentication method
	// Options can be found at https://github.com/moltin/gomo/blob/master/options.go
	client := gomo.NewClient(gomo.ImplicitCredentials(clientID))

	// Execute the debug option function for the client in order to turn on debugging
	client.EnableDebug()

	// Authenticate against the Moltin API
	err := client.Authenticate()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Declare a variable to hold the product including flow data
	var productWithFlowData productFlowData

	err = client.Get(
		"/products/9c40f78c-e35a-41ca-8357-d6d805696371",
		gomo.Data((&productWithFlowData)),
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the body and flow field data of the product returned
	spew.Dump(productWithFlowData)
}
