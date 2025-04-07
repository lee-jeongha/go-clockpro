# go-clockpro
the CLOCK-Pro cache eviction algorithm

> [!NOTE]
> * Original source code: https://github.com/example/example.git
> * Documentation (godoc) for the original source: http://godoc.org/github.com/dgryski/go-clockpro

### Run
```
$ go run clockpro.go main.go --file ./fileio.strace.csv --cachesize 942
```

### Test
```
$ go test clockpro.go clockpro_test.go
```