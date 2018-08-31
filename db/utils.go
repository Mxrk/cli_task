package db

import (
	"encoding/binary"
	"fmt"

	"github.com/boltdb/bolt"
)

// Add a task into a bucket
func AddTask(value string) error {
	err := db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("Tasks"))
		id, _ := b.NextSequence()
		newID := int(id)
		err := b.Put(itob(newID), []byte(value))
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

// Delete a task
func DelTask() {

}

// Get all tasks in one bucket
func GetTasks() error {
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("Tasks"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			id := binary.BigEndian.Uint64(k)
			fmt.Printf("Task number %v: %s\n", id, v)
		}

		return nil
	})
	return nil
}

// Convert int to byte
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
