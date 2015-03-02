package main

import (
	"flag"
	"fmt"
	pb "github.com/kentaro/grpc-gdbm/gdbm"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var port int
var key string
var value string

func init() {
	flag.IntVar(&port, "port", 50051, "port number")
	flag.StringVar(&key, "key", "key", "key name")
	flag.StringVar(&value, "value", "value", "value for key")
	flag.Parse()
}

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGdbmClient(conn)

	r, err := c.Replace(context.Background(), &pb.Request{Key: key, Value: value})
	if err != nil {
		log.Fatalf("gdbm error: %v", err)
	}

	r, err = c.Fetch(context.Background(), &pb.Request{Key: key})
	log.Printf("value for %s: %s", key, r.Value)
}
