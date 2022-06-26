package db

type DB interface {
	ListBuckets() []string
	ListKeys(bucket string) []string
	Get(bucket, key string) string
	PrefixScan(bucket, prefix string, offset, limit int) []string
	PrefixSearchScan(bucket, regex string, offset, limit int) []string
}
