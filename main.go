package main

import (
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

// Version contains the tigs version.
const Version = "0.6.0"

var (
	app = kingpin.New("tigs", "The HTTP client code generator.\n\nSee https://github.com/fgrosse/tigs for further information.")

	pkg       = app.Flag("package", "The name of the package the generated type should be defined in").Required().String()
	name      = app.Flag("name", "The name of the generated HTTP client type.").Default("").String()
	inputFile = app.Flag("in", "The input yaml file to generate the client from").Required().File()
	inputType = app.Flag("type", "The input type").Default("guzzle-yaml").Enum("guzzle-yaml", "guzzle-json")
	verbose   = app.Flag("debug", "Print debug output").Default("false").Bool()
)

func main() {
	app.Version(Version)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	c := &client{
		Package: *pkg,
		Name:    *name,
	}

	err := newDecoder(*inputType, *inputFile).decode(c)
	if err != nil {
		log.Fatal(err)
	}

	logDebug("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	err = generate(os.Stdout, c)
	if err != nil {
		log.Fatal(err)
	}
}
