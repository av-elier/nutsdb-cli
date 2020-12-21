# Test and Debug Notes

### Create a database and insert sample data for testing and debugging.
```shell
$ go run cli.go --db ./test.db debug
```

### One of the test examples for debugging
```shell
$ echo "help
list
list bucket1
get bucket1 k1
get bucketUnknown k2" | go run cli.go --db ./test.db
----------
Commands:
  clear      clear the screen
  exit       exit the program
  get        get <bucket> <key>: show value of given key
  help       display help
  list       `list`: list buckets, or `list <bucket>`: list bucket keys


bucket1
k1
k2
k3
k4
key with spaces
umbrella

```

### in debug configuration it's nice to use `--command=list something`
```shell
$ go run cli.go --db ./test.db --command='get bucket1 k1'
----------
> get bucket1 k1
umbrella
```
