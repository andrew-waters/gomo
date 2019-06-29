package main

import (
	"flag"
	"fmt"

	"github.com/moltin/gomo"
	"github.com/moltin/gomo/core"
)

func main() {
	limitPtr := flag.Int("limit", 100, "maximum page size")
	flag.Parse()

	client := gomo.NewClient()
	err := client.Authenticate()
	if err != nil {
		fmt.Println(err)
		return
	}

	gomo.Iterate(
		*limitPtr,
		func(paginate gomo.RequestResource) error {
			var orders []core.Order
			var meta core.Meta
			err = client.Get(
				"/orders",
				paginate,
				gomo.Data(&orders),
				gomo.Meta(&meta),
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
		},
	)
}
