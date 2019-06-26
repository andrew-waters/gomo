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

	// Fetch orders, two per page
	gomo.Iterate(2, func(paginate gomo.RequestResource, meta *core.Meta) error {
		var orders []core.Order
		_, err = client.Get(
			"/orders",
			paginate,
			gomo.Data(&orders),
		)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf(
			"Page %d/%d, Offset %d, Limit %d, Results %d\n",
			meta.Page.Current,
			meta.Page.Total,
			meta.Page.Offset,
			meta.Page.Limit,
			len(orders),
		)
		return nil
	})
}
