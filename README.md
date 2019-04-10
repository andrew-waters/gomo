# gomo - a Go Client for the Moltin API

[![GoDoc](https://godoc.org/github.com/andrew-waters/gomo?status.svg)](https://godoc.org/github.com/andrew-waters/gomo)
[![Report Card](https://goreportcard.com/badge/github.com/andrew-waters/gomo)](https://goreportcard.com/report/github.com/andrew-waters/gomo)
[![Maintainability](https://api.codeclimate.com/v1/badges/48415c0b8f48979b40a9/maintainability)](https://codeclimate.com/github/andrew-waters/gomo/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/48415c0b8f48979b40a9/test_coverage)](https://codeclimate.com/github/andrew-waters/gomo/test_coverage)

A Golang client for the [moltin](https://moltin.com) API.

```bash
go get github.com/andrew-waters/gomo
```


## Documentation

Code documentation is available on [GoDoc](https://godoc.org/github.com/andrew-waters/gomo).

A guide is avilable on [the wiki](https://github.com/andrew-waters/gomo/wiki), but the demo below shows you how to create a client and get a product with a known ID:

```golang
client := gomo.NewClient(
	gomo.NewClientCredentials(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
	),
)

if err := client.Authenticate(); err != nil {
	log.Fatal(err)
}

products := []entities.Product{}
_, err := client.Get("products", &products)
if err != nil {
	log.Fatal(err)
}

log.Printf("Found %d products\n", len(products))
```


## Testing

In order to fully test the package, you will need a Moltin account to add your credentials to an environment file:

```bash
cp example.env .env
```

Add your credentials and run:

```bash
source .env && go test ./...
```

If you do not suply a client_id and client_secret, we will skip tests that leverage the live API.
