package comm

import (
	"errors"
	"os"
)

type Communicator struct {
	reader *Reader
	writer *Writer
}

func NewCommunicator() *Communicator {
	return &Communicator{
		reader: &Reader{},
		writer: &Writer{
			out: os.Stdout,
		},
	}
}

func (comm *Communicator) Run() error {
	for {
		cmd, err := comm.reader.Read()
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
			// TODO: use ListBuckets() []string
		} else {
			// TODO: use ListKeys(bucket string) []string
		}
	case "get":
		// TODO: use Get(bucket, key string) []string
	default:
		comm.writer.Error("process", errors.New("all parsable commands should be supported, developers error"))
	}
	return nil
}
