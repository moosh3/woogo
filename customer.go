package woogo

import (
	"context"
	"fmt"
	"time"
)

// CustomerURL is the API endpoint that handles requests for a
// customer on WooCommerce
var CustomerURL = "/wp-json/wc/v2/customers"

// CustomerService provides methods to interact with customers on WooCommerce
type CustomerService service

// Customer is a struct representation of the JSON payload recieved
// when calling the WooCommerce API for a customer object
type Customer struct {
	ID               *int      `json:"id,omitempty"`
	DateCreated      time.Time `json:"date_created,omitempty"`
	DateCreatedGMT   time.Time `json:"date_created_gmt,omitempty"`
	DateModified     time.Time `json:"date_modified,omitempty"`
	DateModifiedGMT  time.Time `json:"date_modified_gmt,omitempty"`
	Email            *string   `json:"email,omitempty"`
	FirstName        *string   `json:"first_name,omitempty"`
	LastName         *string   `json:"last_name,omitempty"`
	Role             *string   `json:"role,omitempty"`
	Username         *string   `json:"username,omitempty"`
	Password         *string   `json:"password,omitempty"`
	Billing          Billing   `json:"billing,omitempty"`
	Shipping         Shipping  `json:"shipping,omitempty"`
	IsPayingCustomer *bool     `json:"is_paying_customer,omitempty"`
	OrdersCount      *int      `json:"orders_count,omitempty"`
	TotalSpent       *string   `json:"total_spent,omitempty"`
	AvatarURL        *string   `json:"avatar_url,omitempty"`
	MetaData         MetaData  `json:"meta_data,omitempty"`
}

// Billing is a struct that holds the JSON listed in a Customer object
// in regards to their billing information
type Billing struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Company   *string `json:"company,omitempty"`
	Address1  *string `json:"address_1,omitempty"`
	Address2  *string `json:"address_2,omitempty"`
	City      *string `json:"city,omitempty"`
	State     *string `json:"state,omitempty"`
	PostCode  *string `json:"post_code,omitempty"`
	Country   *string `json:"country,omitempty"`
	Email     *string `json:"email,omitempty"`
	Phone     *string `json:"phone,omitempty"`
}

// Shipping is a struct that holds the JSON listed in a Customer object
// in regards to their shipping information
type Shipping struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Company   *string `json:"company,omitempty"`
	Address1  *string `json:"address_1,omitempty"`
	Address2  *string `json:"address_2,omitempty"`
	City      *string `json:"city,omitempty"`
	State     *string `json:"state,omitempty"`
	PostCode  *string `json:"post_code,omitempty"`
	Country   *string `json:"country,omitempty"`
}

// MetaData is a struct that handles any metadata that is attached
// to the payload recieved when retrieving a Customer object
type MetaData struct {
	ID    int     `json:"id"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// GetByID returns a Customer in a struct representation via
// retrieval using the customers ID
func (cs *CustomerService) GetByID(int id) *Response {
	c := fmt.Sprintf("customers/%d", id)
	req, err := ps.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	product := new(Customer)
	resp, err := ps.client.Do(ctx, req, product)
	if err != nil {
		return nil, resp, err
	}

	return customer, resp, nil
}

// Get takes a string representation of a customer and retrieves it from
// WooCommerce, unmarshalling the response into a new Customer struct
func (cs *CustomerService) Get(ctx context.Context, customer string) *Response {
	var c string
	if customer != "" {
		p = fmt.Sprintf("customers/%v", product)
	} else {
		p = "customer"
	}
	req, err := ps.client.NewRequest("GET", p, nil)
	if err != nil {
		return nil, nil, err
	}

	// Create an empty Customer struct to unmarshall into
	cResp := new(Customer)
	resp, err := ps.client.Do(ctx, req, cResp)
	if err != nil {
		return nil, resp, err
	}

	return cResp, resp, nil
}
