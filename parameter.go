package tigs

// Name represents some input argument to a HTTP endpoint.
type Parameter struct {
	// Name is the name of the argument.
	// Unless specified otherwise, Name is used to transmit this argument over the wire.
	Name string

	// Description is an optional textual representation of this parameter.
	// It is used to generate documentation and is not meant to be used automatically.
	Description string

	// Type represents the type this argument is of.
	// If the type of a parameter is left empty it will be `interface{}` in the generated code.
	// Valid values are:
	//     string
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
	//      json      (TODO)
	Location string

	// Required determines whether this parameter is mandatory or optional.
	// TODO: Implement code to check if required field is there
	Required bool
}

// GeneratedType returns the go type as creating during code generation.
// TODO write individual test
func (p Parameter) GeneratedType() string {
	switch p.Type {
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

// StringCode returns valid go code that will transform the value of this parameter into a string.
// TODO write individual test
func (p Parameter) StringCode() string {
	switch p.Type {
	case "int":
		fallthrough
	case "integer":
		fallthrough
	case "int32":
		fallthrough
	case "int64":
		return `fmt.Sprintf("%d", `+p.Name+`)`
	case "float":
		fallthrough
	case "float32":
		fallthrough
	case "float64":
		return `fmt.Sprintf("%f", `+p.Name+`)`
	case "bool":
		fallthrough
	case "boolean":
		return `fmt.Sprintf("%t", `+p.Name+`)`
	case "":
		return `fmt.Sprintf("%s", `+p.Name+`)`
	default:
		return p.Name
	}
}
