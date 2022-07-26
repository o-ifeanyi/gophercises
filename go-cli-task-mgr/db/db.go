package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

// Initializes bolt database and creates a task bucket if it doesnt exist
func Init(tb, cb []byte, dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(tb)
		_, err = tx.CreateBucketIfNotExists(cb)
		return err
	})
	return err
}

// Adds a new taks to the task bucket
func CreateTask(tb []byte, task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		bct := tx.Bucket(tb)
		id64, _ := bct.NextSequence()
		id = int(id64)
		key := itob(id)
		return bct.Put(key, []byte(task))
	})

	if err != nil {
		return -1, err
	}
	return id, nil
}

// Return a []Task of all task in the task bucket
func AllTasks(tb []byte) ([]Task, error) {
	var allTask []Task
	err := db.View(func(tx *bolt.Tx) error {
		bct := tx.Bucket(tb)
		return bct.ForEach(func(k, v []byte) error {
			allTask = append(allTask, Task{Key: btoi(k), Value: string(v)})
			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return allTask, nil
}

// Delete a task from the task bucket
func DeleteTask(tb []byte, id int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bct := tx.Bucket(tb)
		key := itob(id)
		return bct.Delete(key)
	})

	if err != nil {
		return err
	}
	return nil
}

// Resets the entire task list and db
func ResetTask(tb []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(tb)
	})

	if err != nil {
		return err
	}
	return nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
