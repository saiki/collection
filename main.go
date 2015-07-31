package main

import (
	log "github.com/Sirupsen/logrus"
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "saiki"
	app.Email = ""
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{})

	app.Run(os.Args)
}
