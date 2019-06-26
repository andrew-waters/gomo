package main

import (
	"fmt"

	"github.com/moltin/gomo"
	"github.com/moltin/gomo/core"
)

func main() {
	client := gomo.NewClient()
	err := client.Authenticate()
	if err != nil {
		fmt.Println(err)
		return
	}

	var meta core.Meta
	_, err = client.Get("/orders", gomo.Meta(&meta))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Total orders: %d\n", meta.Results.Total)
}
