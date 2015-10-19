package tigs

import (
	"io"
	"fmt"
	"strings"
)

type formattableWriter struct {
	io.Writer
}

func (w *formattableWriter) printf(format string, a ...interface{}) (n int, err error) {
	format = strings.Replace(format, "    ", "\t", -1) + "\n"
	return fmt.Fprintf(w, format, a...)
}
