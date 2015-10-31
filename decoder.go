package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type decoder struct {
	unmarshaller
	io.Reader
}

var registeredUnmarshallers = map[string]unmarshaller{}

type unmarshaller interface {
	Unmarshal(input []byte, c *client) (err error)
}

func newDecoder(inputType string, input io.Reader) (decoder, error) {
	u, isDefined := registeredUnmarshallers[inputType]
	if !isDefined {
		return decoder{}, fmt.Errorf("unknown input type %q", inputType)
	}

	return decoder{u, input}, nil
}

func (d decoder) decode(c *client) error {
	input, err := ioutil.ReadAll(d)
	if err != nil {
		return err
	}

	return d.Unmarshal(input, c)
}

func sanitizeYAML(input []byte) []byte {
	var sanitizedInput = &bytes.Buffer{}

	line := &bytes.Buffer{}
	lineBeginning := true
	for _, c := range input {
		switch c {
		case '\n':
			if strings.TrimSpace(line.String()) != "" {
				sanitizedInput.Write(append(line.Bytes(), '\n'))
				line.Reset()
				lineBeginning = true
			}
		case '\t':
			if lineBeginning {
				line.WriteString("    ")
			} else {
				line.WriteByte(c)
			}
		case ' ':
			line.WriteByte(c)
		default:
			lineBeginning = false
			line.WriteByte(c)
		}
	}

	sanitizedInput.Write(line.Bytes())

	s := sanitizedInput.Bytes()
	return s
}
