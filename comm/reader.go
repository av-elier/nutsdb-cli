package comm

import (
	"bufio"
	"os"
)

type Reader struct{}

type Cmd struct {
	t string
}

var CmdEnd = Cmd{t: "EOF"}

func (r *Reader) Read() (Cmd, error) {
	s := bufio.NewScanner(os.Stdin)
	ok := s.Scan()
	if !ok {
		return CmdEnd, nil
	}

	text := s.Text()
	// TODO: parse and validate input

	return Cmd{t: text}, nil
}
