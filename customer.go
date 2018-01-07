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
