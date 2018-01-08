package woogo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
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

// A Client establishes and interacts with WooCommerce's HTTP API
type Client struct {
	// HTTP client used for making requests
	Client *http.Client
	// Base WooCommerce URL (your store location)
	BaseURL *url.URL

	common service

	Customers *CustomerService
	Products  *ProductService

	// ConsumerKey
	ck string
	// ConsumerSecret
	cs string
	// Version
	v string
	// WP_API
	wp bool
	// QueryStringAUth
	q string
}

// service Reuse a single struct instead of allocating one for each service on the heap.
type service struct {
	client *Client
}

// RawType represents type of raw format of a request instead of JSON.
type RawType uint8

const (
	// Diff format.
	Diff RawType = 1 + iota
	// Patch format.
	Patch
)

// RawOptions specifies parameters when user wants to get raw format of
// a response instead of JSON.
type RawOptions struct {
	Type RawType
}

// NewClient returns a new instance of the Client struct
func NewClient(h *http.Client) *Client {
	if h == nil {
		h = http.DefaultClient
	}
	c := &Client{Client: h}
	c.common.client = c
	c.Customers = (*CustomerService)(&c.common)
	c.Products = (*ProductService)(&c.common)
	return c
}

// NewRequest handles requests
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	signature, p := OAuthSignature(c.ck, c.cs, method, path)
	url := parseURL(path, signature, p) // p.Add(signature)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

// Response wraps a *http.Response object
type Response struct {
	*http.Response
	// Handle pagination
	NextPage int
	PrvPage  int
}

// Do is passed requests from NewRequest and executes it, returning a wrapped *http.Response
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = withContext(ctx, req)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	// Wrap *http.Response and return it for validation, etc
	response := c.newResponse(resp)
	// Custom validation on a *Response struct
	err = c.CheckResponse(response)
	if err != nil {
		// Return anyway for debugging
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				// Ignore errors from empty response
				err = nil
			}
		}
	}

	return response, err
}

// newResponse is used internally to resturn a Response object
// containing a pointer to a response and paginate options
func (c *Client) newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// CheckResponse is a validator that takes a
func (c *Client) CheckResponse(r *Response) error {
	if r.Body != nil {
		return nil
	}

	return nil
}

// withContext returns a context that can be passed though API calls
func withContext(ctx context.Context, r *http.Request) *http.Request {
	return r
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
