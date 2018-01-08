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

// var authURL = "?" + "oauth_consumer_key=" + p.Get("ck") + "&oauth_nonce=" + p.Get("nonce") + "&oauth_signature=" + url.QueryEscape(signature) + "&oauth_signature_method=" + p.Get("sig_method") + "&oauth_timestamp=" + p.Get("timestamp")

func parseURL(path, signature string, params *url.Values) string {
	return ""
}

// OAuthSignature returns a signature and oauth parameters
func OAuthSignature(ck, cs, method, reqURL string) (string, *url.Values) {
	hash := "HMAC-SHA256"

	p := url.Values{} // URL params
	p.Add("ck", ck)
	p.Add("nonce", newNonce())
	p.Add("sig_method", hash)
	p.Add("timestamp", strconv.Itoa(int(time.Now().Unix())))

	url := strings.Join([]string{strings.ToUpper(method), url.QueryEscape(reqURL), url.QueryEscape(p.Encode())}, "&")

	return newSignature(cs, url), &p
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
