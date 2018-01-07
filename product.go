package woogo

import (
		"time"

		"github.com/buger/jsonparser"
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


data := []byte{
  ```{
	  "id": 794,
	  "name": "Premium Quality",
	  "slug": "premium-quality-19",
	  "permalink": "https://example.com/product/premium-quality-19/",
	  "date_created": "2017-03-23T17:01:14",
	  "date_created_gmt": "2017-03-23T20:01:14",
	  "date_modified": "2017-03-23T17:01:14",
	  "date_modified_gmt": "2017-03-23T20:01:14",
	  "type": "simple",
	  "status": "publish",
	  "featured": false,
	  "catalog_visibility": "visible",
	  "description": "<p>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.</p>\n",
	  "short_description": "<p>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.</p>\n",
	  "sku": "",
	  "price": "21.99",
	  "regular_price": "21.99",
	  "sale_price": "",
	  "date_on_sale_from": null,
	  "date_on_sale_from_gmt": null,
	  "date_on_sale_to": null,
	  "date_on_sale_to_gmt": null,
	  "price_html": "<span class=\"woocommerce-Price-amount amount\"><span class=\"woocommerce-Price-currencySymbol\">&#36;</span>21.99</span>",
	  "on_sale": false,
	  "purchasable": true,
	  "total_sales": 0,
	  "virtual": false,
	  "downloadable": false,
	  "downloads": [],
	  "download_limit": -1,
	  "download_expiry": -1,
	  "external_url": "",
	  "button_text": "",
	  "tax_status": "taxable",
	  "tax_class": "",
	  "manage_stock": false,
	  "stock_quantity": null,
	  "in_stock": true,
	  "backorders": "no",
	  "backorders_allowed": false,
	  "backordered": false,
	  "sold_individually": false,
	  "weight": "",
	  "dimensions": {
	    "length": "",
	    "width": "",
	    "height": ""
	  },
	  "shipping_required": true,
	  "shipping_taxable": true,
	  "shipping_class": "",
	  "shipping_class_id": 0,
	  "reviews_allowed": true,
	  "average_rating": "0.00",
	  "rating_count": 0,
	  "related_ids": [
	    53,
	    40,
	    56,
	    479,
	    99
	  ],
	  "upsell_ids": [],
	  "cross_sell_ids": [],
	  "parent_id": 0,
	  "purchase_note": "",
	  "categories": [
	    {
	      "id": 9,
	      "name": "Clothing",
	      "slug": "clothing"
	    },
	    {
	      "id": 14,
	      "name": "T-shirts",
	      "slug": "t-shirts"
	    }
	  ],
	  "tags": [],
	  "images": [
	    {
	      "id": 792,
	      "date_created": "2017-03-23T14:01:13",
	      "date_created_gmt": "2017-03-23T20:01:13",
	      "date_modified": "2017-03-23T14:01:13",
	      "date_modified_gmt": "2017-03-23T20:01:13",
	      "src": "https://example.com/wp-content/uploads/2017/03/T_2_front-4.jpg",
	      "name": "",
	      "alt": "",
	      "position": 0
	    },
	    {
	      "id": 793,
	      "date_created": "2017-03-23T14:01:14",
	      "date_created_gmt": "2017-03-23T20:01:14",
	      "date_modified": "2017-03-23T14:01:14",
	      "date_modified_gmt": "2017-03-23T20:01:14",
	      "src": "https://example.com/wp-content/uploads/2017/03/T_2_back-2.jpg",
	      "name": "",
	      "alt": "",
	      "position": 1
	    }
	  ],
	  "attributes": [],
	  "default_attributes": [],
	  "variations": [],
	  "grouped_products": [],
	  "menu_order": 0,
	  "meta_data": [],
	  "_links": {
	    "self": [
	      {
	        "href": "https://example.com/wp-json/wc/v2/products/794"
	      }
	    ],
	    "collection": [
	      {
	        "href": "https://example.com/wp-json/wc/v2/products"
	      }
	    ]
	  }
	}
```
}
