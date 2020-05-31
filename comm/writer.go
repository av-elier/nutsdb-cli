package comm

import (
	"fmt"
	"io"
)

type Writer struct {
	out io.Writer
}

// TODO: function with table formatting?

func (w *Writer) WriteString(s string) (int, error) {
	return w.out.Write([]byte(s + "\n"))
}

func (w *Writer) WritePromt() (int, error) {
	return w.out.Write([]byte("> "))
}

func (w *Writer) WriteStrings(ss []string) (int, error) {
	written := 0
	var err error = nil
	for _, s := range ss {
		l, err := fmt.Fprintf(w.out, "%s\n", s)
		written += l
		if err != nil {
			return written, err
		}
	}
	return written, err
}

func (w *Writer) Error(stage string, err error) (int, error) {
	return fmt.Fprintf(w.out, "error in %s: %s\n\n", stage, err)
}
