package command

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
	"github.com/codegangsta/cli"
)

// CmdAdd: add text.
func CmdAdd(c *cli.Context) {
	path, err := storePath(c.GlobalString("path"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Add.")
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		log.Debugln("close db.")
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	tx, err := db.Begin(true)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range c.Args() {
		log.Debugln(v)
	}
}
