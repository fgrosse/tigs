package tigs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type FooClient struct {
	BaseURL *url.URL
	Client *http.Client
}

func NewFooClient(baseURL string) (*FooClient, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL for new FooClient: %s", err)
	}

	return &FooClient{
		BaseURL: u,
		Client: http.DefaultClient,
	}, nil
}

func (c *FooClient) GetStuff(s string, b bool, i int, i32 int32, i64 int64, f32 float32, f64 float64, foo interface{}) (*http.Response, error) {
	u, err := c.BaseURL.Parse("/stuff")
	if err != nil {
		return nil, err
	}

	u.Query().Add("s", s)
	u.Query().Add("b", fmt.Sprintf("%t", b))
	u.Query().Add("i", fmt.Sprintf("%d", i))
	u.Query().Add("i32", fmt.Sprintf("%d", i32))
	u.Query().Add("i64", fmt.Sprintf("%d", i64))
	u.Query().Add("f32", fmt.Sprintf("%f", f32))
	u.Query().Add("f64", fmt.Sprintf("%f", f64))
	u.Query().Add("foo", fmt.Sprintf("%s", foo))

	return c.Client.Get(u.String())
}

func (c *FooClient) FancyURI(s string, b bool, i int, i32 int32, i64 int64, f32 float32, f64 float64, foo interface{}) (*http.Response, error) {
	// TODO replace "/stuff/{category}/{id}/"

	u, err := c.BaseURL.Parse("/stuff")
	if err != nil {
		return nil, err
	}

	return c.Client.Get(u.String())
}

func (c *FooClient) PostStuff(s string, b bool, i int, foo interface{}) (*http.Response, error) {
	u, err := c.BaseURL.Parse("/stuff")
	if err != nil {
		return nil, err
	}

	u.Query().Add("s", s)

	data, err := json.Marshal(map[string]interface{}{
		"b":   b,
		"i":   i,
		"foo": foo,
	})

	if err != nil {
		return nil, fmt.Errorf("could not marshal body for PostStuff: %s", err)
	}

	return c.Client.Post(u.String(), "application/json", bytes.NewBuffer(data))
}

func (c *FooClient) PostForm(s string, b bool, i int, i32 int32, i64 int64, f32 float32, f64 float64, foo interface{}) (*http.Response, error) {
	u, err := c.BaseURL.Parse("/stuff")
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Add("s", s)
	v.Add("b", fmt.Sprintf("%t", b))
	v.Add("i", fmt.Sprintf("%d", i))
	v.Add("i32", fmt.Sprintf("%d", i32))
	v.Add("i64", fmt.Sprintf("%d", i64))
	v.Add("f32", fmt.Sprintf("%f", f32))
	v.Add("f64", fmt.Sprintf("%f", f64))
	v.Add("foo", fmt.Sprintf("%s", foo))

	return c.Client.PostForm(u.String(), v)
}
