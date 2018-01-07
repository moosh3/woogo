# woogo

[![GoDoc](https://godoc.org/github.com/aleccunningham/woogo?status.svg)](https://godoc.org/github.com/aleccunningham/woogo)

![Imgur](https://i.imgur.com/5KVjCyft.jpg)

A WooCommerce API in Golang

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

`export DOMAIN=`
`export CONSUMER_KEY=`
`export CONSUMER_SECRET=`
