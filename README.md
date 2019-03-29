# gomo - a Moltin Go Client

A Golang client for the [moltin](https://moltin.com) API.

`go get github.com/andrew-waters/gomo`

## Instantiating a Client

### Client Credentials

THe most common type of authentication for secure execution environments (eg a server you control).

```golang
client, err := gomo.NewClient(
  gomo.NewClientCredentials(
    ":client_id",
    ":client_secret",
  ),
)

if err != nil {
  log.Fatal(err)
}
```

### Implicit

This should be used if your code could be exposed. If you only need read access to your data, this is recommended.

```golang
client, err := gomo.NewClient(
  gomo.NewImplicitCredentials(
    ":client_id",
  ),
)

if err != nil {
  log.Fatal(err)
}
```

### Using custom endpoints

You can instatiate a client with a custom API endpoint using the `NewClientWithCustomEndpoint` func:

```golang
client, err := gomo.NewClientWithCustomEndpoint(
  gomo.NewClientCredentials(
    ":client_id",
    ":client_secret",
  ),
  "https://alt.domain.tld",
)
```

## Working with Resources

API calls are made on the client instance using the verb as a func name (eg `client.Get("products", &products)`).

Because Go is statically typed, you need to provide a pointer to a concrete struct that you want to map the API response to and optionally an include pointer for included resources.

To get resources from the API:

```golang
product := gomo.Product{} // this is a standard Moltin product
wrapper, err := client.Get("products/:id", &product)
```

In this example, `product` will now contain the product info from the response ready for you to use and if you need to work with the entire response, for example getting included items, you can use the returned `response` - `response.Included`.

The reason that we pass in the `moltin.Product{}` param is that when working with Moltin, you can use flows to extend resources in the API which does not suit a statically typed language.

Therefore, if your product contained custom fields you can create your own struct and pass it in. This allows you to get all the benefits of the language with the flexibility of the API.

```golang
type YourProduct struct {
  ID              string `json:"id"`
  CustomAttribute string `json:"custom_attribute"`
}
product := YourProduct{} // this is your custom product struct
wrapper, err := client.Get("products/:id", &product)
```

To include data, you should pass your target included object where the response will be unmarshalled to:

```golang
product := entities.Product{}
included  := entities.ProductIncludes{}
wrapper, err := client.Get("products/:id", &product, &included)
```


## Examples

Examples are included in the `./examples` directory. To use them you should copy `example.env` to `.env` and replace the Client ID and Secret values with the ones from your own store.

From that directory you can run an example using `go run auth/main.go`, `go run products/read/main.go`.
