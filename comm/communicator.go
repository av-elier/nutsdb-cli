package comm

type Communicator struct {
	reader    *Reader
	writer    *Writer
	processor *Processor
}

func NewCommunicator() *Communicator {
	return &Communicator{
		reader:    &Reader{},
		processor: &Processor{},
		writer:    &Writer{},
	}
}

func (comm *Communicator) Run() error {
	for {
		cmd, err := comm.reader.Read()
		if err != nil {
			comm.writer.Error("read", err)
			continue
		}
		out, err := comm.processor.Process(cmd)
		if err != nil {
			comm.writer.Error("process", err)
			continue
		}
		comm.writer.Strings(out)
	}
}
