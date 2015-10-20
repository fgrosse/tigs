package tigs

// Client holds all information necessary to generate a go client for an HTTP web service.
type Client struct {
	// Name is the name of the generated go type for this client.
	Name string

	// Package is the name of the package that the generated client code will be defined in.
	Package string

	// Endpoints is a list of HTTP endpoints that are available over this client.
	Endpoints []Endpoint
}

func (c Client) ContainsJSONEndpoints() bool {
	for _, ep := range c.Endpoints {
		if ep.HasJSONParameter() {
			return true
		}
	}

	return false
}
