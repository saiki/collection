package command

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
	"github.com/codegangsta/cli"
	"strconv"
)

// CmdAdd: add text.
func CmdAdd(c *cli.Context) {
	path, err := storePath(c.GlobalString("path"))
	if err != nil {
		log.Fatal(err)
	}
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
	if err != nil {
		log.Fatalln(err)
	}
	err = persist(db, c.Args())
	if err != nil {
		log.Fatalln(err)
	}
}

func persist(db *bolt.DB, args []string) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	log.Debugln("start transaction.")
	for _, v := range args {
		log.Debugln(v)
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				return err
			}
			count := btoi(bucket.Get([]byte(v)))
			count += 1
			err = bucket.Put([]byte(v), itob(count))
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	tx.Commit()
	return nil
}

func btoi(b []byte) int {
	if b == nil {
		return 0
	}
	str := string(b[:bytes.Index(b, []byte{0})])
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}

func itob(i int) []byte {
	return []byte(strconv.Itoa(i))
}
