package tigs

// An Endpoint represents an operation of a service that can be accessed by clients via an URL.
type Endpoint struct {
	// Name is the symbolic name of this endpoint.
	// It is used in code generate and does not represent any part of the URL that is actually used.
	// Every endpoint must always have a Name or it is considered invalid.
	Name string

	// Description is an optional textual representation of this endpoint.
	// It is used to generate documentation and is not meant to be used automatically.
	Description string

	// Method is the HTTP method of this endpoint.
	Method string

	// URL is the URL under which the endpoint can be reached.
	URL string

	// Parameters is the list of parameters that are used to create the HTTP request.
	Parameters []Parameter

	// client is set by the generator before the corresponding endpoint function is generated.
	// This only serves as utility function and to make some method functions less complicated.
	client Client
}
