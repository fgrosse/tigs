package main

import (
	"bytes"
	"text/template"
)

func loadTemplate(name string) *template.Template {
	resource, err := Asset(name)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	lineBeginning := true
	spacesRead := 0
	for _, c := range resource {
		switch {
		case c == '\n':
			lineBeginning = true
			spacesRead = 0
			buf.WriteByte(c)
		case lineBeginning && c == ' ':
			if spacesRead == 3 { // three plus this iteration makes four
				spacesRead = 0
				buf.WriteString("\t")
			} else {
				spacesRead++
			}
		default:
			lineBeginning = false
			buf.WriteByte(c)
		}
	}

	tmpl := template.New(name)
	tmpl.Parse(buf.String())

	return tmpl
}
