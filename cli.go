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
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "db",
			Usage:       "path to nutsdb batabase",
			Destination: &args.db,
		},
	}
	app.Action = func(c *cli.Context) error {
		var db db.DB // NewDB(args.db)
		comm := comm.Communicator{}

		fmt.Println("hello nutsdb", db, comm)
		return comm.Run()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
