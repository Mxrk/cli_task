package db

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func Init() error {
	var err error
	db, err = bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal("Could not create database", err)
	}
	//	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}
