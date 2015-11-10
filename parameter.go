package main

import "fmt"

// parameter represents some input argument to a HTTP endpoint.
type parameter struct {
	// Name is the name of the argument.
	// Unless specified otherwise, Name is used to transmit this argument over the wire.
	Name string

	// Description is an optional textual representation of this parameter.
	// It is used to generate documentation and is not meant to be used automatically.
	Description string

	// Type represents the type this argument is of.
	// If the type of a parameter is left empty it will be `interface{}` in the generated code.
	// Valid values are:
	//     string|text
	//     int|int32|int64|integer|int
	//     float|float32|float64
	//     bool|boolean
	Type string

	// Location determines where this argument appears in the request.
	// If the location is left empty it will default to `query`
	// Valid values are:
	//      query
	//      uri       (TODO)
	//      postField (TODO)
	//      json
	Location string

	// Required determines whether this parameter is mandatory or optional.
	Required bool
}

// generatedType returns the go type as creating during code generation.
func (p parameter) generatedType() string {
	switch p.Type {
	case "text":
		return "string"
	case "integer":
		return "int"
	case "float":
		return "float64"
	case "boolean":
		return "bool"
	case "":
		return "interface{}"
	default:
		return p.Type
	}
}

// stringCode returns valid go code that will transform the value of this parameter into a string.
func (p parameter) stringCode() string {
	switch p.Type {
	case "string":
		fallthrough
	case "text":
		return p.Name
	case "int":
		fallthrough
	case "integer":
		fallthrough
	case "int32":
		fallthrough
	case "int64":
		return `fmt.Sprintf("%d", ` + p.Name + `)`
	case "float":
		fallthrough
	case "float32":
		fallthrough
	case "float64":
		return `fmt.Sprintf("%f", ` + p.Name + `)`
	case "bool":
		fallthrough
	case "boolean":
		return `fmt.Sprintf("%t", ` + p.Name + `)`
	case "":
		fallthrough
	default:
		return `fmt.Sprintf("%s", ` + p.Name + `)`
	}
}

func (p parameter) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("missing name")
	}

	if p.Type == "" {
		return fmt.Errorf("missing type")
	}

	switch p.Location {
	case "query":
	case "json":
	case "uri":
	case "postField":
		// the above are all valid
	case "":
		return fmt.Errorf("missing location")
	default:
		return fmt.Errorf("unknown location %q", p.Location)
	}

	return nil
}
