package comm

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type Reader struct {
	in io.Reader
}

type Cmd struct {
	t      string
	bucket string
	key    string
}

var emptyInputErr = errors.New("empty input")
var unknownCommandInputErr = errors.New("unknown command")

var CmdEnd = Cmd{t: "exit"}

// TODO: Support buckets with space in it

func (r *Reader) Read() (Cmd, error) {
	s := bufio.NewScanner(r.in)
	ok := s.Scan()
	if !ok {
		return CmdEnd, nil
	}

	text := s.Text()
	spaceSplitted := strings.Split(text, " ")
	if len(spaceSplitted) == 0 || spaceSplitted[0] == "" {
		return Cmd{}, emptyInputErr
	}
	res := Cmd{}
	res.t = spaceSplitted[0]

	switch res.t {
	case "list":
		if len(spaceSplitted) >= 2 {
			res.bucket = strings.Join(spaceSplitted[1:], " ")
		}
	case "get":
		if len(spaceSplitted) < 3 {
			return res, errors.New("usage: get <bucket> <key>")
		}
		res.bucket = spaceSplitted[1]
		res.key = strings.Join(spaceSplitted[2:], " ")
	case "help":
	case "exit":
	default:
		return res, unknownCommandInputErr
	}

	return res, nil
}
