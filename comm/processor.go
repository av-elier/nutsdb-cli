package comm

type Processor struct{}

// TODO: return not []string but a type with one of many bodies

func (r *Processor) Process(cmd Cmd) ([]string, error) {
	return []string{string(cmd.t)}, nil
}
