package main

import (
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

// Version contains the tigs version.
const Version = "0.4.0"

var (
	app = kingpin.New("tigs", "The HTTP client code generator.\n\nSee https://github.com/fgrosse/tigs for further information.")

	inputFile = app.Flag("in", "The input yaml file to generate the client from").Required().File()
	inputType = app.Flag("type", "The input type").Default("guzzle-yaml").Enum("guzzle-yaml", "guzzle-json")
)

func main() {
	app.Version(Version)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	var c client
	err := newDecoder(*inputType, *inputFile).decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	err = generate(os.Stdout, c)
	if err != nil {
		log.Fatal(err)
	}
}
