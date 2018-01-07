package woogo

import (
	"time"
)

var (
	ProductPOSTRequest    = "/wp-json/wc/v2/customers"
	ProductGETListRequest = "/wp-json/wc/v2/customers"
	// Attribute "<id>" required
	ProductGETDetailRequest = "/wp-json/wc/v2/customers/%s"
	ProductPUTRequest       = "/wp-json/wc/v2/customers/%s"
	ProductDELETERequest    = "/wp-json/wc/v2/customers/%s"
)

type Product struct {
	ID                int
	Name              string
	Slug              string
	Permalink         string
	DateCreated       time.Time
	DateCreatedGMT    time.Time
	DateModified      time.Time
	DateModifiedGMT   time.Time
	Type              string
	Status            string
	Featured          bool
	CatalogVisibility string
	Description       string
	ShortDescription  string
	SKU               string
	Price             string
	RegularPrice      string
	SalePrice         string
	DateOnSaleFrom    time.Time
	DateOnSaleFromGMT time.Time
	DateOnSaleTo      time.Time
	DateOnSaleToGMT   time.Time
	PriceHTML         string
	OnSale            bool
	Purchasable       bool
	TotalSales        int
	Virtual           bool
	Downloadable      bool
	Downloads         []byte // List of downloadable files
	DownloadLimit     int
	DownloadExpiry    int
	ExternalURL       string
	ButtonText        string
	TaxStatus         string
	TaxClass          string
	ManageStock       bool
	StockQuanityt     int
	InStock           bool
	BackOrders        string
	BackOrdersAllowed bool
	BackOrdered       bool
	SoldIndividually  bool
	Weight            string
	Dimensions        []byte // Product dimensions
	ShippingRequired  bool
	ShippingTaxable   bool
	ShippingClass     string
	ShippingClassID   string
	ReviewsAllowed    bool
	AverageRating     string
	RatingCount       int
	RelatedIDs        []byte // List of related product IDs
	UpsellIDs         []byte // List of up-sell producct IDs
	CrossSellIDs      []byte // list of cross-sell product IDs
	ParentID          int
	PurchaseNote      string
	Categories        []byte // List of categories
	Tags              []byte // List of tags
	Images            []byte // List of images
	Attributes        []byte // List of attributes
	DefaultAttributes []byte // Defaults variation attributes
	Variations        []byte // List of variations IDs
	GroupedProducts   []byte // List of grouped products ID
	MenuOrder         int
	MetaData          []byte // Meta data
}

func (woo *Client) ListByID(int id) {
	/* ... */
}
