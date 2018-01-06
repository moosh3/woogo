package woogo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	// DomainPath WooCommerce store domain
	DomainPath = os.Getenv("domain")
	// ConsumerKey WooCommerce API Consumer key
	ConsumerKey = os.Getenv("consumer_key")
	// ConsumerSecret WooCommerce API Consumer secret
	ConsumerSecret = os.Getenv("consumer_secret")
	// Version API Version
	Version = "wc/v2/"
	// QueryStringAuth API Auth string
	QueryStringAuth = true
	// WpAPI API connection
	WpAPI = true
)

// WooCommerceService defines an interfaces for interacting with
// the WooCommerce API
type WooCommerceService interface {
	Get(Path string, Params *url.Values) io.ReadCloser
	Post(Path string, Data interface{}) io.ReadCloser
	Put(Path string, Data interface{}) io.ReadCloser
	Delete(Path string, Params *url.Values) io.ReadCloser
	Patch(Path string, Params *url.Values) io.ReadCloser
}

// Client is a struct that executes methods
// in the WooCommerceService
type Client struct {
	Client *http.Client
	Domain *url.URL
	ck     string // ConsumerKey
	cs     string // ConsumerSecret
	v      string // Version
	wp     bool   // WP_API
	q      string // QueryStringAUth
}

// NewClient returns a new instance of the Client struct
func NewClient(domainURL, ck, cs string) *Client {
	domain, err := url.Parse(domainURL)
	if err != nil {
		panic(err)
	}
	domain.Path = "/wp-json/wc/v2/"
	c := &Client{Client: &http.Client{},
		Domain: domain, ck: ck, cs: cs,
	}
	return c
}

// Data is a payload sent to WC
var Data interface{}

// handler handles requests
func (woo *Client) handler(Method, Path string, Params *url.Values, Data interface{}) io.ReadCloser {
	var method = strings.ToUpper(Method)

	signature, p := OAuthSignature(woo.ck, woo.cs, method, Path)
	p.Add("oauth_signature", signature)

	url := woo.Domain.String() + Path + "?" + "oauth_consumer_key=" + p.Get("ck") + "&oauth_nonce=" + p.Get("nonce") + "&oauth_signature=" + url.QueryEscape(signature) + "&oauth_signature_method=" + p.Get("sig_method") + "&oauth_timestamp=" + p.Get("timestamp")

	body := new(bytes.Buffer)
	encoder := json.NewEncoder(body)
	if err := encoder.Encode(Data); err != nil {
		return nil
	}

	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("user-agent", "WooCommerce API Client-Go\n")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := woo.Client.Do(req)
	if err != nil {
		return nil
	}
	return resp.Body
}

// Get gets
func (woo *Client) Get(Path string, Params *url.Values) io.ReadCloser {
	if Params != nil {
		return woo.handler("GET", Path, Params, nil)
	}
	return woo.handler("GET", Path, nil, nil)
}

// Post posts
func (woo *Client) Post(Path string, Data interface{}) io.ReadCloser {
	if Data != nil {
		return woo.handler("POST", Path, nil, Data)
	}
	return woo.handler("POST", Path, nil, nil)
}

// Put puts
func (woo *Client) Put(Path string, Data interface{}) io.ReadCloser {
	if Data != nil {
		return woo.handler("PUT", Path, nil, Data)
	}
	return woo.handler("PUT", Path, nil, nil)
}

// Delete deletes
func (woo *Client) Delete(Path string, Params *url.Values) io.ReadCloser {
	if Params != nil {
		return woo.handler("DELETE", Path, Params, nil)
	}
	return woo.handler("DELETE", Path, nil, nil)
}

// Patch patches
func (woo *Client) Patch(Path string, Params *url.Values) io.ReadCloser {
	if Params != nil {
		return woo.handler("PATCH", Path, Params, nil)
	}
	return woo.handler("PATCH", Path, nil, nil)
}
