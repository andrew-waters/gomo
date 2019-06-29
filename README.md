# gomo - a Go Client for the Moltin API

[![GoDoc](https://godoc.org/github.com/moltin/gomo?status.svg)](https://godoc.org/github.com/moltin/gomo)
[![Report Card](https://goreportcard.com/badge/github.com/moltin/gomo)](https://goreportcard.com/report/github.com/moltin/gomo)
[![Maintainability](https://api.codeclimate.com/v1/badges/48415c0b8f48979b40a9/maintainability)](https://codeclimate.com/github/moltin/gomo/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/48415c0b8f48979b40a9/test_coverage)](https://codeclimate.com/github/moltin/gomo/test_coverage)

A Golang client for the [moltin](https://moltin.com) API.

```bash
go get github.com/moltin/gomo
```


## Documentation

Reference documentation is available on
[GoDoc](https://godoc.org/github.com/moltin/gomo).

```golang
client := gomo.NewClient(
	gomo.ClientCredentials(
		os.Getenv("MOLTIN_CLIENT_ID"),
		os.Getenv("MOLTIN_CLIENT_SECRET"),
	),
)

if err := client.Authenticate(); err != nil {
	log.Fatal(err)
}

products := []entities.Product{}
err := client.Get("products", gomo.Data(&products))
if err != nil {
	log.Fatal(err)
}

log.Printf("Found %d products\n", len(products))
```


## Testing

In order to fully test the package, you will need a Moltin account to add your
credentials to an environment file:

```bash
cp .env.example .env
```

Add your credentials and run:

```bash
source .env && go test ./...
```

If you do not supply a `MOLTIN_CLIENT_ID` and `MOLTIN_CLIENT_SECRET`, we will
skip tests that leverage the live API.
