package woogo

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetCustomers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	customers := make([]map[string]interface{}, 0)

	// Mock GET request to customers
	httpmock.RegisterResponder("GET", "https://test.com/articles.json",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, customers)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
}

func TestPostCustomers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	customers := make([]map[string]interface{}, 0)

	// Mock POST request to customers
	httpmock.RegisterResponder("POST", "https://api.mybiz.com/customers.json",
		func(req *http.Request) (*http.Response, error) {
			customer := make(map[string]interface{})
			if err := json.NewDecoder(req.Body).Decode(&customer); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}

			customers = append(customers, customer)

			resp, err := httpmock.NewJsonResponse(200, customer)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
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
