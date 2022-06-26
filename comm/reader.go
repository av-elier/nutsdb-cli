package comm

import (
	"errors"
	"strings"
)

type Reader struct {
}

type Cmd struct {
	t      string
	bucket string
	key    string
	prefix string
	regex  string
}

var errUnknownCommandInput = errors.New("unknown command")

func (r *Reader) Read(t string, args []string) (Cmd, error) {
	res := Cmd{}

	switch t {
	case "list":
		if len(args) >= 1 {
			res.bucket = strings.Join(args, " ")
		}
	case "prefix":
		if len(args) < 1 {
			return res, errors.New("usage: prefix <bucket> <prefix>, bucket and/or prefix not specified")
		}
		res.bucket = args[0]
		res.prefix = strings.Join(args[1:], " ")
	case "regex":
		if len(args) < 1 {
			return res, errors.New("usage: regex <bucket> <regex>, bucket and/or regex not specified")
		}
		res.bucket = args[0]
		res.regex = strings.Join(args[1:], " ")
	case "get":
		if len(args) < 2 {
			return res, errors.New("usage: get <bucket> <key>, bucket and/or key not specified")
		}
		res.bucket = args[0]
		res.key = strings.Join(args[1:], " ")
	default:
		return res, errUnknownCommandInput
	}

	return res, nil
}
