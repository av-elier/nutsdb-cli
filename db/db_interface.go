package db

import (
	"fmt"
	"log"

	"github.com/xujiajun/nutsdb"
)

type DB interface {
	ListBuckets() []string
	ListKeys(bucket string) []string
	// Get(bucket, key string) []string
}

type Db struct {
	DbDir string
}

// CreateDebugDb just creates some non-empty database for debug purposes
func (file Db) CreateDebugDb() {
	opt := nutsdb.DefaultOptions
	opt.Dir = file.DbDir
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal("Cannot open database")
	}

	bucket := "bucket1"
	keysAndValues := []struct {
		Key   []byte
		Value []byte
	}{
		{[]byte("k1"), []byte("umbrella")},
		{[]byte("k2"), []byte("budger")},
	}

	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			for _, kv := range keysAndValues {
				if err := tx.Put(bucket, kv.Key, kv.Value, 0); err != nil {
					log.Fatalln(err)
					return err
				}
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

// ListBuckets returns all buckets in given database
func (file Db) ListBuckets() []string {
	opt := nutsdb.DefaultOptions
	opt.Dir = file.DbDir
	db, err := nutsdb.Open(opt)
	defer db.Close()
	if err != nil {
		return []string{}
	}
	buckets := []string{}
	for k := range db.BPTreeIdx {
		buckets = append(buckets, k)
	}
	return buckets
}

// ListKeys returns all keys in given bucket
func (file Db) ListKeys(bucket string) []string {
	opt := nutsdb.DefaultOptions
	opt.Dir = file.DbDir
	db, err := nutsdb.Open(opt)
	defer db.Close()

	if err != nil {
		return nil
	}

	buckets := []string{}

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			prefix := []byte{}
			// Constrain 100 entries returned
			// bucketBin := []byte(bucket)
			limit := 1000 // TODO
			if entries, err := tx.PrefixScan(bucket, prefix, limit); err != nil {
				log.Println("Error occured while scanning for keys", err)
				return nil
			} else {
				for _, entry := range entries {
					keyAsString := string(entry.Key)
					buckets := append(buckets, keyAsString)
					fmt.Println(buckets)
				}
			}
			return nil
		}); err != nil {
		log.Println(err)
	}

	return buckets
}
