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

  req := client.Get("products")
  var body []byte
  decoder := json.NewDecoder(req)
  err := decoder.Decode(&body)
  if err != nil {
    panic(err)
  }
  defer body.Close()

  fmt.Printf("Product list: %s", body.Text)
}
