package tigshttp

import (
	"net/http"
	"net/url"
)

// NewRequest creates a new request from a given HTTP method and url.
// Note that neither the body nor the ContentLength will be set.
//
// The difference to http.NewRequest is that this function can be used
// in case you already have a URL and want to avoid marshalling it
// back into a string and doing error related handling again.
func NewRequest(method string, u *url.URL) *http.Request {
	return &http.Request{
		Method:     method,
		URL:        u,
		Host:       u.Host,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
	}
}
