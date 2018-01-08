package woogo

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"
)

// CustomerURL is the API endpoint that handles requests for a
// customer on WooCommerce
var ProductURL = "/wp-json/wc/v2/products"

// ProductService provides methods to interact with products on WooCommerce
type ProductService service

// Product is a struct representation of the JSON payload recieved
// when calling the WooCommerce API for a product object
type Product struct {
	ID                *int       `json:"id,omitempty"`
	Name              *string    `json:"name,omitempty"`
	Slug              *string    `json:"slug,omitempty"`
	Permalink         *string    `json:"permalink,omitempty"`
	DateCreated       time.Time `json:"date_created,omitempty"`
	DateCreatedGMT    time.Time `json:"date_created_gmt,omitempty"`
	DateModified      time.Time `json:"date_modified,omitempty"`
	DateModifiedGMT   time.Time `json:"date_modified_gmt,omitempty"`
	Type              *string    `json:"type,omitempty"`
	Status            *string    `json:"status,omitempty"`
	Featured          *bool      `json:"featured,omitempty"`
	CatalogVisibility *string    `json:"catalog_visibility,omitempty"`
	Description       *string    `json:"description,omitempty"`
	ShortDescription  *string    `json:"short_description,omitempty"`
	SKU               *string    `json:"sku,omitempty"`
	Price             *string    `json:"price,omitempty"`
	RegularPrice      *string    `json:"regular_price,omitempty"`
	SalePrice         *string    `json:"sale_price,omitempty"`
	DateOnSaleFrom    time.Time `json:"date_on_sale_from,omitempty"`
	DateOnSaleFromGMT time.Time `json:"date_on_sale_from_gmt,omitempty"`
	DateOnSaleTo      time.Time `json:"date_on_sale_to,omitempty"`
	DateOnSaleToGMT   time.Time `json:"date_on_sale_to_gmt,omitempty"`
	PriceHTML         *string    `json:"price_html,omitempty"`
	OnSale            *bool      `json:"on_sale,omitempty"`
	Purchasable       *bool      `json:"purchasable,omitempty"`
	TotalSales        *int       `json:"total_sales,omitempty"`
	Virtual           *bool      `json:"virtual,omitempty"`
	Downloadable      *bool      `json:"downloadable,omitempty"`
	Downloads         []byte    // List of downloadable files
	DownloadLimit     *int       `json:"download_limit,omitempty"`
	DownloadExpiry    *int       `json:"download_expiry,omitempty"`
	ExternalURL       *string    `json:"external_url,omitempty"`
	ButtonText        *string    `json:"button_text,omitempty"`
	TaxStatus         *string    `json:"tax_status,omitempty"`
	TaxClass          *string    `json:"tax_class,omitempty"`
	ManageStock       *bool      `json:"manage_stock,omitempty"`
	StockQuanity      *int       `json:"stock_quantity,omitempty"`
	InStock           *bool      `json:"in_stock,omitempty"`
	BackOrders        *string    `json:"back_orders,omitempty"`
	BackOrdersAllowed *bool      `json:"back_orders_allowed,omitempty"`
	BackOrdered       *bool      `json:"back_ordered,omitempty"`
	SoldIndividually  *bool      `json:"sold_individually,omitempty"`
	Weight            *string    `json:"weight,omitempty"`
	Dimensions        []byte    // Product dimensions
	ShippingRequired  *bool      `json:"shipping_required,omitempty"`
	ShippingTaxable   *bool      `json:"shipping_taxable,omitempty"`
	ShippingClass     *string    `json:"shipping_class,omitempty"`
	ShippingClassID   *string    `json:"shipping_class_id,omitempty"`
	ReviewsAllowed    *bool      `json:"reviews_allowed,omitempty"`
	AverageRating     *string    `json:"average_rating,omitempty"`
	RatingCount       *int       `json:"rating_count,omitempty"`
	RelatedIDs        []byte    // List of related product IDs
	UpsellIDs         []byte    // List of up-sell producct IDs
	CrossSellIDs      []byte    // list of cross-sell product IDs
	ParentID          *int       `json:"parent_id,omitempty"`
	PurchaseNote      *string    `json:"purchase_note,omitempty"`
	Categories        []byte    // List of categories
	Tags              []byte    // List of tags
	Images            []byte    // List of images
	Attributes        []byte    // List of attributes
	DefaultAttributes []byte    // Defaults variation attributes
	Variations        []byte    // List of variations IDs
	GroupedProducts   []byte    // List of grouped products ID
	MenuOrder         *int       `json:"menu_order,omitempty"`
	MetaData          []byte    // Meta data
}

// Get takes a product id and retrieves it from
// WooCommerce, unmarshalling the response into a new Product struct
func (ps *ProductService) GetByID(ctx context.Context, int id) *Response {
	if id != ""
	p := fmt.Sprintf("products/%d", id)
	req, err := ps.client.NewRequest("GET", p, nil)
	if err != nil {
		return nil, nil, err
	}

	product := new(Product)
	resp, err := ps.client.Do(ctx, req, product)
	if err != nil {
		return nil, resp, err
	}

	return product, resp, nil
}

// Get takes a string representation of a product and retrieves it from
// WooCommerce, unmarshalling the response into a new Product struct
func (ps *ProductService) Get(ctx context.Context, product string) *Response {
	var p string
	if product != "" {
		p = fmt.Sprintf("products/%v", product)
	} else {
		p = "product"
	}
	req, err := ps.client.NewRequest("GET", p, nil)
	if err != nil {
		return nil, nil, err
	}

	// Create an empty Product struct to unmarshall into
	pResp := new(Product)
	resp, err := ps.client.Do(ctx, req, pResp)
	if err != nil {
		return nil, resp, err
	}

	return pResp, resp, nil
}
