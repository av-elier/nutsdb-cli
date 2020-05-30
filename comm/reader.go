package comm

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Reader struct{}

type Cmd struct {
	t      string
	bucket string
	key    string
}

var emptyInputErr = errors.New("empty input")
var unknownCommandInputErr = errors.New("unknown command")

var CmdEnd = Cmd{t: "exit"}

func (r *Reader) Read() (Cmd, error) {
	s := bufio.NewScanner(os.Stdin)
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
	case "get":
		switch len(spaceSplitted) {
		case 1:
			return res, errors.New("get <bucket> [<key>]: bucket is required")
		case 3:
			res.bucket = spaceSplitted[1]
			res.key = strings.Join(spaceSplitted[2:], " ")
		case 2:
			res.bucket = spaceSplitted[1]
		}
	case "help":
	case "exit":
	default:
		return res, unknownCommandInputErr
	}

	return res, nil
}
