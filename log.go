package main

import (
	"fmt"
	"os"
)

func logDebug(message string, args ...interface{}) {
	if *verbose == false {
		return
	}

	logMessage(message, args...)
}

func logMessage(message string, args ...interface{}) {
	writer := os.Stdout
	//if *outputPath == "" {
	// since we already output the generated code on stdout we print messages on stderr
	writer = os.Stderr
	//}

	fmt.Fprintf(writer, message+"\n", args...)
}
