package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

// Version contains the tigsgen version.
const Version = "0.2.0"

var (
	app = kingpin.New("tigs", "The HTTP client code generator.\n\nSee https://github.com/fgrosse/tigs for further information.")
)

func main() {
	defer panicHandler()
	app.Version(Version)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	/// TODO parse client from input
	c := ServiceClient{
		Name:    "TestClient",
		Package: "tigs",
		Endpoints: []Endpoint{
			{
				Name:        "GetThings",
				Description: "GetThings fetches things for you!!",
				Method:      "GET",
				URL:         "/things",
			},
			{
				Name:        "DoStuff",
				Description: "DoStuff does cool stuff",
				Method:      "POST",
				URL:         "/do/stuff",
				Parameters: []Parameter{
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
			{
				Name:   "PatchThis",
				Method: "PATCH",
				URL:    "/my/awesome/patch",
				Parameters: []Parameter{
					{
						Name:     "param",
						Type:     "float",
						Location: "json",
					},
				},
			},
		},
	}

	err := Generate(os.Stdout, c)
	if err != nil {
		log.Fatal(err)
	}
}

func panicHandler() {
	if r := recover(); r != nil {
		log.Fatalf("FATAL ERROR: %s", r)
	}
}
