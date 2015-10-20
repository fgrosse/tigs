package main

import (
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
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

	c := exampleClient // TODO parse client from input instead

	err := generate(os.Stdout, c)
	if err != nil {
		log.Fatal(err)
	}
}

var exampleClient = client{
	Name:    "TestClient",
	Package: "tigs",
	Endpoints: []endpoint{
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
			Parameters: []parameter{
				{
					name:        "param1",
					description: "The first parameter",
					typeString:  "string",
					location:    "json",
					required:    true,
				},
				{
					name:        "param2",
					description: "The second parameter",
					location:    "json",
					required:    false,
				},
			},
		},
		{
			Name:   "PatchThis",
			Method: "PATCH",
			URL:    "/my/awesome/patch",
			Parameters: []parameter{
				{
					name:       "param",
					typeString: "float",
					location:   "json",
				},
			},
		},
	},
}

func panicHandler() {
	if r := recover(); r != nil {
		log.Fatalf("FATAL ERROR: %s", r)
	}
}
