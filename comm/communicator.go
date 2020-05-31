package comm

import (
	"errors"
	"io"
	"os"

	"github.com/av-elier/nutsdb-cli/db"
)

type Communicator struct {
	reader *Reader
	writer *Writer
	db     db.DB
}

func NewCommunicator(in io.Reader, db db.DB) *Communicator {
	return &Communicator{
		reader: &Reader{in: in},
		writer: &Writer{
			out: os.Stdout,
		},
		db: db,
	}
}

func (comm *Communicator) Run() error {
	for {
		comm.writer.WritePromt()
		cmd, err := comm.reader.Read()
		if err == emptyInputErr {
			continue
		}
		if err != nil {
			comm.writer.Error("read", err)
			continue
		}
		if cmd.t == "exit" { // the only flow control cmd
			return nil
		}
		err = comm.processCmd(cmd)
		if err != nil {
			comm.writer.Error("process", err)
		}
	}
}

func (comm *Communicator) processCmd(cmd Cmd) error {
	switch cmd.t {
	case "help":
		comm.writer.WriteString(`available commands:
	help
		display this message
	list
		list buckets
	list <bucket>
		list bucket keys
	get <bucket> <key>
		show value of given key
	exit
		exit

	buckets with space in it, umm... sorry`)
	case "list":
		if cmd.bucket == "" {
			list := comm.db.ListBuckets()
			comm.writer.WriteStrings(list)
		} else {
			list := comm.db.ListKeys(cmd.bucket)
			comm.writer.WriteStrings(list)
		}
	case "get":
		value := comm.db.Get(cmd.bucket, cmd.key)
		comm.writer.WriteString(value)
	default:
		comm.writer.Error("process", errors.New("all parsable commands should be supported, developers error"))
	}
	return nil
}
