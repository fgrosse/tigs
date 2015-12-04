package main

import (
	"log"
	"os"

	"io/ioutil"

	"gopkg.in/alecthomas/kingpin.v2"
)

// Version contains the tigs version.
const Version = "0.9.2"

var (
	app = kingpin.New("tigs", "The HTTP client code generator.\n\nSee https://github.com/fgrosse/tigs for further information.")

	pkg         = app.Flag("package", "The name of the package the generated type should be defined in").Required().String()
	Name        = app.Flag("name", "The name of the generated HTTP client type.").Default("").String()
	inputFile   = app.Flag("in", "The input file to generate the client from").Required().File()
	inputType   = app.Flag("type", "The input type").Default("guzzle-yaml").Enum("guzzle-yaml", "guzzle-json")
	verbose     = app.Flag("debug", "Print debug output").Default("false").Bool()
	forceExport = app.Flag("forceExport", "Always export all client functions even if their names are lower case in the input file").Bool()

	Debug = log.New(ioutil.Discard, "Debug: ", 0)
	Error = log.New(os.Stderr, "Error: ", 0)
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("Error: ")
}

func main() {
	app.Version(Version)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *verbose {
		EnableDebugLog()
	}

	c := &client{
		Package: *pkg,
		Name:    *Name,
	}

	Debug.Printf("Decoding input from file %q", (*inputFile).Name())
	s := settings{inheritance: true, forceExport: *forceExport}
	Debug.Printf("Settings: %+v", s)

	err := newDecoder(*inputType, *inputFile).decode(c, s)
	if err != nil {
		Error.Fatal(err)
	}

	Debug.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	err = generate(os.Stdout, c)
	if err != nil {
		Error.Fatal(err)
	}
}

func EnableDebugLog() {
	Debug = log.New(os.Stderr, "Debug: ", 0)
}
