// The tigsgen binary
package main

import (
	"os"

	"github.com/fgrosse/tigs"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

// Version contains the tigsgen version.
const Version = "0.1.0"

var (
	app = kingpin.New("tigsgen", "The HTTP client code generator.\n\nSee https://github.com/fgrosse/tigs for further information.")
)

func main() {
	defer panicHandler()
	app.Version(Version)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	/// TODO parse client from input
	c := tigs.Client{
		Name:    "TestClient",
		Package: "foobar",
		Endpoints: []tigs.Endpoint{
			{
				Name:        "DoStuff",
				Description: "DoStuff does cool stuff",
				Method:      "GET",
				URL:         "/do/stuff",
				Parameters: []tigs.Parameter{
					{
						Name:        "param1",
						Description: "The first parameter",
						Type:        "string",
						Location:    "json",
						Required:    true,
					},
					{
						Name:        "param2",
						Description: "The second parameter",
						Location:    "json",
						Required:    false,
					},
				},
			},
		},
	}

	err := tigs.Generate(os.Stdout, c)
	if err != nil {
		log.Fatal(err)
	}
}

func panicHandler() {
	if r := recover(); r != nil {
		log.Fatalf("FATAL ERROR: %s", r)
	}
}
