package database

import (
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

var dbName = "test.db"

func SetDbName(str string) {
	dbName = str
}

func Init() {
	db, err := bolt.Open(dbName, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket([]byte("film"))
		tx.CreateBucket([]byte("person"))
		tx.CreateBucket([]byte("planet"))
		tx.CreateBucket([]byte("species"))
		tx.CreateBucket([]byte("starship"))
		tx.CreateBucket([]byte("vehicle"))

		return nil
	}); err != nil {
		log.Fatal(err)
	}
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

func DeleteDB() {
	os.Remove(dbName)
}

func Update(bucketName []byte, key []byte, value []byte) {
	db, err := bolt.Open(dbName, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := db.Update(func(tx *bolt.Tx) error {
		// Create a bucket.
		if err := tx.Bucket(bucketName).Put(key, value); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

func GetValue(bucketName []byte, key []byte) string {
	db, err := bolt.Open(dbName, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	var result []byte
	if err := db.View(func(tx *bolt.Tx) error {
		//value = tx.Bucket([]byte(bucketName)).Get(key)
		byteLen := len(tx.Bucket([]byte(bucketName)).Get(key))
		result = make([]byte, byteLen)
		copy(result[:], tx.Bucket([]byte(bucketName)).Get(key)[:])
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}

	return string(result)
}

//debug
func CheckBucket(bucketName []byte) {
	db, err := bolt.Open(dbName, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			log.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
