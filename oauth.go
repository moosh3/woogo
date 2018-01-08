package woogo

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var authURL = "?" + "oauth_consumer_key=" + p.Get("ck") + "&oauth_nonce=" + p.Get("nonce") + "&oauth_signature=" + url.QueryEscape(signature) + "&oauth_signature_method=" + p.Get("sig_method") + "&oauth_timestamp=" + p.Get("timestamp")

func parseURL(path, signature, params) string {
	return nil
}

// OAuthSignature returns a signature and oauth parameters
func OAuthSignature(ConsumerKey, ConsumerSecret, Method, RequestURL string) (string, *url.Values) {
	hash := "HMAC-SHA256"

	params := url.Values{}
	params.Add("ck", ConsumerKey)
	params.Add("nonce", newNonce())
	params.Add("sig_method", hash)
	params.Add("timestamp", strconv.Itoa(int(time.Now().Unix())))

	url := strings.Join([]string{strings.ToUpper(Method), url.QueryEscape(RequestURL), url.QueryEscape(params.Encode())}, "&")

	return newSignature(ConsumerSecret, url), &params
}

func newNonce() string {
	nonce := make([]byte, 16)
	rand.Read(nonce)
	sha := fmt.Sprintf("%x", sha1.Sum(nonce))
	return sha
}

func newSignature(ConsumerSecret, URL string) string {
	mac := hmac.New(sha256.New, []byte(ConsumerSecret+"&"))
	mac.Write([]byte(URL))
	bytes := mac.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(bytes)
	return signature
}
