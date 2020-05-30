# Test and Debug Notes

```shell
go run cli.go debug

echo "help
list
list bucket1
get bucket1 k1
get bucketUnknown k2" | go run cli.go
```

in debug configuration it's nice to use `--command=list something`
