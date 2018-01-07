# woogo

![Imgur](https://i.imgur.com/5KVjCyft.jpg)

A WooCommerce API in Golang! Example usage:

```Go
package main

import (
  "fmt"

  "github.com/aleccunningham/woogo"
)

func main() {
  var domain = "woocommerce.com"
  client := woogo.NewClient(domain)

  patient, err := client.Get("products", profileID)
  err := jsonparser.EachKey(func() {
          /* ... */
         })
  fmt.Printf("Product list: %s", body.Text)
}
