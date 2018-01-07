package woogo

import "time"

var (
	CustomerPOSTRequest    = "/wp-json/wc/v2/customers"
	CustomerGETListRequest = "/wp-json/wc/v2/customers"
	// Attribute "<id>" required
	CustomerGETDetailRequest = "/wp-json/wc/v2/customers/%s"
	CustomerPUTRequest       = "/wp-json/wc/v2/customers/%s"
	CustomerDELETERequest    = "/wp-json/wc/v2/customers/%s"
)

type Customer struct {
	ID               int
	DateCreated      time.Time
	DateCreatedGMT   time.Time
	DateModified     time.Time
	DateModifiedGMT  time.Time
	Email            string
	FirstName        string
	LastName         string
	Role             string
	Username         string
	Password         string
	Billing          Billing
	Shipping         Shipping
	IsPayingCustomer bool
	OrdersCount      int
	TotalSpent       string
	AvatarURL        string
	MetaData         MetaData
}

type Billing struct {
	FirstName string
	LastName  string
	Company   string
	Address1  string
	Address2  string
	City      string
	State     string
	PostCode  string
	Country   string
	Email     string
	Phone     string
}

type Shipping struct {
	FirstName string
	LastName  string
	Company   string
	Address1  string
	Address2  string
	City      string
	State     string
	PostCode  string
	Country   string
}

type MetaData struct {
	ID    int
	Key   string
	Value string
}

func (woo *Client) ListByID(int id) {
	/* ... */
}

/*
GETSample := []byte{
  ```{
    "id": 25,
    "date_created": "2017-03-21T16:09:28",
    "date_created_gmt": "2017-03-21T19:09:28",
    "date_modified": "2017-03-21T16:09:30",
    "date_modified_gmt": "2017-03-21T19:09:30",
    "email": "john.doe@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "role": "customer",
    "username": "john.doe",
    "billing": {
      "first_name": "John",
      "last_name": "Doe",
      "company": "",
      "address_1": "969 Market",
      "address_2": "",
      "city": "San Francisco",
      "state": "CA",
      "postcode": "94103",
      "country": "US",
      "email": "john.doe@example.com",
      "phone": "(555) 555-5555"
    },
    "shipping": {
      "first_name": "John",
      "last_name": "Doe",
      "company": "",
      "address_1": "969 Market",
      "address_2": "",
      "city": "San Francisco",
      "state": "CA",
      "postcode": "94103",
      "country": "US"
    },
    "is_paying_customer": false,
    "orders_count": 0,
    "total_spent": "0.00",
    "avatar_url": "https://secure.gravatar.com/avatar/8eb1b522f60d11fa897de1dc6351b7e8?s=96",
    "meta_data": [],
    "_links": {
      "self": [
        {
          "href": "https://example.com/wp-json/wc/v2/customers/25"
        }
      ],
      "collection": [
        {
          "href": "https://example.com/wp-json/wc/v2/customers"
        }
      ]
    }
  }
```
}
*/
