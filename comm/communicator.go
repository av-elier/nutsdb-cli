package comm

import (
	"io"

	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
	"github.com/av-elier/nutsdb-cli/db"
)

type Communicator struct {
	db     db.DB
	reader *Reader
	shell  *ishell.Shell
}

func NewCommunicator(in io.ReadCloser, db db.DB) *Communicator {
	comm := &Communicator{
		db:     db,
		reader: &Reader{},
		shell: ishell.NewWithConfig(&readline.Config{
			Prompt: "> ",
			Stdin:  in,
		}),
	}
	comm.shell.EOF(func(c *ishell.Context) {
		comm.shell.Stop()
	})
	comm.shell.SetHomeHistoryPath(".nutsdb_cli_history")

	comm.shell.AddCmd(&ishell.Cmd{
		Name: "list",
		Help: "`list`: list buckets, or `list <bucket>`: list bucket keys",
		Completer: func(args []string) []string {
			if len(args) != 0 {
				return nil
			}
			return comm.db.ListBuckets()
		},
		Func: func(c *ishell.Context) {
			nutsCmd, err := comm.reader.Read(c.Cmd.Name, c.Args)
			if err != nil {
				c.Err(err)
				return
			}
			if nutsCmd.bucket == "" {
				list := comm.db.ListBuckets()
				comm.printLines(list)
			} else {
				list := comm.db.ListKeys(nutsCmd.bucket)
				comm.printLines(list)
			}
		},
	})
	comm.shell.AddCmd(&ishell.Cmd{
		Name: "get",
		Help: "get <bucket> <key>: show value of given key",
		Completer: func(args []string) []string {
			if len(args) == 0 {
				return comm.db.ListBuckets()
			}
			if len(args) == 1 {
				return comm.db.ListKeys(args[0])
			}
			return nil
		},
		Func: func(c *ishell.Context) {
			nutsCmd, err := comm.reader.Read(c.Cmd.Name, c.Args)
			if err != nil {
				c.Err(err)
				return
			}
			value := comm.db.Get(nutsCmd.bucket, nutsCmd.key)
			comm.shell.Println(value)

		},
	})
	return comm
}

func (comm *Communicator) Run() error {
	comm.shell.Run()
	return nil
}

func (comm *Communicator) printLines(lines []string) {
	for _, l := range lines {
		comm.shell.Println(l)
	}
}
