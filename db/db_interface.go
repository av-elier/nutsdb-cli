package db

import (
	"log"

	"github.com/xujiajun/nutsdb"
)

type DB interface {
	ListBuckets() []string
	// ListKeys(bucket string) []string
	// Get(bucket, key string) []string
}

type Db struct {
	DbDir string
}

func (file Db) CreateDebugDb() {
	opt := nutsdb.DefaultOptions
	opt.Dir = file.DbDir
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal("Cannot open database")
	}

	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			key := []byte("name1")
			val := []byte("val1")
			bucket := "bucket1"
			if err := tx.Put(bucket, key, val, 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func (file Db) ListBuckets() []string {
	opt := nutsdb.DefaultOptions
	opt.Dir = file.DbDir
	db, err := nutsdb.Open(opt)
	if err != nil {
		return []string{}
	}
	buckets := []string{}
	for k := range db.BPTreeIdx {
		buckets = append(buckets, k)
	}
	defer db.Close()
	return buckets
}
