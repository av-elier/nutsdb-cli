package main

import (
	"bytes"
	"io"
	"io/ioutil"
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
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "command",
				Usage:       "a command to execute instead of starting repl",
				Destination: &args.command,
			},
			cli.StringFlag{
				Name:        "db",
				Usage:       "path to nutsdb batabase",
				Destination: &args.db,
				Required:    true,
			},
		},
		Commands: []cli.Command{
			{
				Name:  "debug",
				Usage: "Create debug database 'test.db' in cwd",
				Action: func(c *cli.Context) error {
					cwd, _ := os.Getwd()
					impl := db.NutsDB{DbDir: path.Join(cwd, args.db)}
					impl.CreateDebugDb()
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			cwd, _ := os.Getwd()
			nuts := db.NutsDB{DbDir: path.Join(cwd, args.db)}
			var in io.ReadCloser = os.Stdin
			if args.command != "" {
				in = ioutil.NopCloser(bytes.NewReader([]byte(args.command)))
			}
			comm := comm.NewCommunicator(in, nuts)
			return comm.Run()
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
