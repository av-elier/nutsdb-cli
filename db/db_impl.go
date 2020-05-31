package db

import (
	"log"

	"github.com/xujiajun/nutsdb"
)

type NutsDB struct {
	DbDir string
}

// CreateDebugDb just creates some non-empty database for debug purposes
func (nuts NutsDB) CreateDebugDb() {
	opt := nutsdb.DefaultOptions
	opt.Dir = nuts.DbDir
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
		{[]byte("k3"), []byte("some text value with spaces")},
		{[]byte("k4"), []byte("some text value with spaces and \nnew \nlines")},
		{[]byte("key with spaces"), []byte("3.1415926")},
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
func (nuts NutsDB) ListBuckets() []string {
	opt := nutsdb.DefaultOptions
	opt.Dir = nuts.DbDir
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
func (nuts NutsDB) ListKeys(bucket string) []string {
	opt := nutsdb.DefaultOptions
	opt.Dir = nuts.DbDir
	db, err := nutsdb.Open(opt)
	defer db.Close()

	if err != nil {
		return nil
	}

	keys := []string{}

	err = db.View(func(tx *nutsdb.Tx) error {
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
				keys = append(keys, keyAsString)
			}
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	return keys
}

func (nuts NutsDB) Get(bucket, key string) string {
	opt := nutsdb.DefaultOptions
	opt.Dir = nuts.DbDir
	db, err := nutsdb.Open(opt)
	defer db.Close()
	if err != nil {
		return ""
	}
	var res string
	err = db.View(func(tx *nutsdb.Tx) error {
		entry, err := tx.Get(bucket, []byte(key))
		if err != nil {
			return err
		}
		res = string(entry.Value)
		return nil
	})
	return res
}
