# gomo - a Moltin Go Client

[![GoDoc](https://godoc.org/github.com/andrew-waters/gomo?status.svg)](https://godoc.org/github.com/andrew-waters/gomo)
[![Maintainability](https://api.codeclimate.com/v1/badges/48415c0b8f48979b40a9/maintainability)](https://codeclimate.com/github/andrew-waters/gomo/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/48415c0b8f48979b40a9/test_coverage)](https://codeclimate.com/github/andrew-waters/gomo/test_coverage)

A Golang client for the [moltin](https://moltin.com) API.

`go get github.com/andrew-waters/gomo`


## Documentation

Code documentation is available on [GoDoc](https://godoc.org/github.com/andrew-waters/gomo).

A guide is avilable on [the wiki](https://github.com/andrew-waters/gomo/wiki).


## Testing

In order to fully test the package, you will need a Moltin account to add your credentials to an environment file:

```
cp example.env .env
```

Add your credntials and run:

```
source .env && go test ./...
```

If you do not suply a client_id and client_secret, we will skip tests that leverage the live API.
