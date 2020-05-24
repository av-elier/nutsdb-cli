package comm

import (
	"fmt"
	"os"
)

type Writer struct{}

// TODO: function with table formatting?

func (w *Writer) Strings(ss []string) {
	for _, s := range ss {
		fmt.Fprintf(os.Stdout, "%s\n", s)
	}
	fmt.Fprintln(os.Stdout)
}

func (w *Writer) Error(stage string, err error) {
	fmt.Fprintf(os.Stderr, "error in stage %s: %s\n", stage, err)
}
