package main

import (
	"fmt"
	"log"
	"os"

	"github.com/av-elier/nutsdb-cli/comm"
	"github.com/av-elier/nutsdb-cli/db"
	"github.com/urfave/cli"
)

type args struct {
	db string
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
					impl := db.Db{DbDir: cwd + "/test.db"}
					impl.CreateDebugDb()
					return nil
				},
			},
		},
	}
	app.Flags = []cli.Flag{
		// TODO: make this argument required
		cli.StringFlag{
			Name:        "db",
			Usage:       "path to nutsdb batabase",
			Destination: &args.db,
		},
	}
	app.Action = func(c *cli.Context) error {
		comm := comm.Communicator{}

		// fmt.Println("hello nutsdb", db, comm)
		fmt.Println("hello nutsdb")
		return comm.Run()
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
