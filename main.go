package main

import (
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

// Version contains the tigs version.
const Version = "0.3.0"

var (
	app = kingpin.New("tigs", "The HTTP client code generator.\n\nSee https://github.com/fgrosse/tigs for further information.")

	inputFile = app.Flag("in", "The input yaml file to generate the client from").Required().File()
)

func main() {
	app.Version(Version)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	d, err := newDecoder("yaml", *inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var c client
	err = d.decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	err = generate(os.Stdout, c)
	if err != nil {
		log.Fatal(err)
	}
}
