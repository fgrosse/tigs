package tigshttp

import (
	"bytes"
	"fmt"
	"net/url"
)

// ExpandURITemplate is a very minimalistic implementation of uri templates.
// It only supports efficiently replacing simple variables.
// Have a look at the tests to see an example.
func ExpandURITemplate(template string, args map[string]interface{}) (*url.URL, error) {
	u := new(bytes.Buffer)
	tmp := new(bytes.Buffer)

	isExpanding := false
	for _, r := range template {
		switch {
		case r == '{' && !isExpanding:
			isExpanding = true
		case r == '}' && isExpanding:
			v := tmp.String()
			replacement, isKnown := args[v]
			if !isKnown {
				replacement = "{" + v + "}"
			}

			u.WriteString(fmt.Sprintf("%v", replacement))
			tmp.Reset()
			isExpanding = false
		case isExpanding:
			tmp.WriteRune(r)
		default:
			u.WriteRune(r)
		}
	}

	if isExpanding {
		return nil, fmt.Errorf("invalid URI template: missing }")
	}

	return url.Parse(u.String())
}
