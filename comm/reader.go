package comm

import (
	"bufio"
	"os"
)

type Reader struct{}

type Cmd string

func (r *Reader) Read() (Cmd, error) {
	stdin := bufio.NewReader(os.Stdin)
	return stdin.ReadString("\n")
}
