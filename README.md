# grpc-gdbm

A network interface to gdbm using [grpc-go](https://github.com/grpc/grpc-go).

(_This is just just an example code of gRPC_)

## Usage

server:

```
$ go run server/main.go
```

client:

```
$ go run client/main.go -key foo -value bar
2015/03/03 01:11:12 value for foo: bar
```

## For Developer

If you want to re-generate go code from IDC file:

```
$ protoc -I ./protos ./protos/gdbm.proto --go_out=plugins=grpc:gdbm
```

## Author

[Kentaro Kuribayashi](http://kentarok.org)
