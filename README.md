# woogo

[![GoDoc](https://godoc.org/github.com/aleccunningham/woogo?status.svg)](https://godoc.org/github.com/aleccunningham/woogo)

A WooCommerce API in Golang. The package documentation can be found on godocs (link above) and the official documentation for the full WooCommerce REST API can be found [here](https://woocommerce.github.io/woocommerce-rest-api-docs/)

## Usage

```Go
import "github.com/aleccunningham/woogo"
```

Construct a new WooCommerce client, and the use the various functions to access data. For example:

```Go
client := woogo.NewClient(ctx, "storename.com")

// list all product info for <product_id>
products, _, err := client.Products.ListByID(ctx, "65", nil)
}
```

## Authentication

`woogo` includes functionality to generate URLs signed via an oauth token to authenticate with the API. You will need to export the following environment variables with their key values from a generated token on WooCommerce's website. Alternatively you can c/p them into the library itself.

```
$ export DOMAIN=
$ export CONSUMER_KEY=
$ export CONSUMER_SECRET=
```
