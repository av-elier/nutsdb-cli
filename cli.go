package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/av-elier/nutsdb-cli/comm"
	"github.com/av-elier/nutsdb-cli/db"
	"github.com/urfave/cli"
)

type args struct {
	db string

	command string
}

func main() {
	var args args
	app := &cli.App{
		Name: "nutsdb-cli",
		Commands: []cli.Command{
			{
				Name:  "debug",
				Usage: "Create debug database 'test.db' in cwd",
				Action: func(c *cli.Context) error {
					cwd, _ := os.Getwd()
					impl := db.NutsDB{DbDir: path.Join(cwd, "test.db")}
					impl.CreateDebugDb()
					return nil
				},
			},
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "command",
			Usage:       "a command to execute instead of starting repl",
			Destination: &args.command,
		},
		// TODO: make this argument required
		cli.StringFlag{
			Name:        "db",
			Usage:       "path to nutsdb batabase",
			Destination: &args.db,
		},
	}
	app.Action = func(c *cli.Context) error {
		cwd, _ := os.Getwd()
		nuts := db.NutsDB{DbDir: path.Join(cwd, "test.db")}
		var in io.Reader = os.Stdin
		if args.command != "" {
			in = bytes.NewReader([]byte(args.command))
		}
		comm := comm.NewCommunicator(in, nuts)
		fmt.Println("hello nutsdb")
		return comm.Run()
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
