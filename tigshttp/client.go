// Package tigshttp provides basic interfaces and utility functions for generated HTTP clients.
package tigshttp

import "net/http"

// A Client sends http.Requests and returns http.Responses or errors in case of
// failure. Responses with StatusCode >= 400 are *not* considered a failure.
// http.Client implements client
type Client interface {
	Do(*http.Request) (*http.Response, error)
}

// ClientFunc is a function type that implements the Client interface.
type ClientFunc func(*http.Request) (*http.Response, error)

// Do implements the Client interface by executing f with the given request as input.
func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

// A decorator wraps a Client with additional behaviour or capabilities.
type Decorator func(Client) Client

// Decorate wraps a Client with all the given Decorators, in order.
func Decorate(c Client, ds ...Decorator) Client {
	for _, d := range ds {
		c = d(c)
	}
	return c
}
