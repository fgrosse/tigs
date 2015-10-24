package main

// parameter represents some input argument to a HTTP endpoint.
type parameter struct {
	// name is the name of the argument.
	// Unless specified otherwise, Name is used to transmit this argument over the wire.
	Name string

	// description is an optional textual representation of this parameter.
	// It is used to generate documentation and is not meant to be used automatically.
	Description string

	// typeString represents the type this argument is of.
	// If the type of a parameter is left empty it will be `interface{}` in the generated code.
	// Valid values are:
	//     string|text
	//     int|int32|int64|integer|int
	//     float|float32|float64
	//     bool|boolean
	TypeString string

	// location determines where this argument appears in the request.
	// If the location is left empty it will default to `query`
	// Valid values are:
	//      query
	//      uri       (TODO)
	//      postField (TODO)
	//      json
	Location string

	// Required determines whether this parameter is mandatory or optional.
	// TODO: Implement code to check if required field is there
	Required bool
}

// generatedType returns the go type as creating during code generation.
func (p parameter) generatedType() string {
	switch p.TypeString {
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
		return p.TypeString
	}
}

// stringCode returns valid go code that will transform the value of this parameter into a string.
func (p parameter) stringCode() string {
	switch p.TypeString {
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
