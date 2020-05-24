package db

type DB interface {
	ListBuckets() []string
	ListKeys(bucket string) []string
	Get(bucket, key string) []string
}
