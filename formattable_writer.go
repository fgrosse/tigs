package main

import (
	"fmt"
	"io"
)

type formattableWriter struct {
	io.Writer
}

func (w *formattableWriter) printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format+"\n", a...)
}
