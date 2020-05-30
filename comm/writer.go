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

func (w *Writer) WriteStrings(ss []string) (int, error) {
	written := 0
	for _, s := range ss {
		l, err := fmt.Fprintf(w.out, "%s\n", s)
		written += l
		if err != nil {
			return written, err
		}
	}
	l, err := fmt.Fprintln(w.out)
	return written + l, err
}

func (w *Writer) Error(stage string, err error) (int, error) {
	return fmt.Fprintf(w.out, "error in %s: %s\n", stage, err)
}
