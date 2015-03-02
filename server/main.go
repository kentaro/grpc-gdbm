package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/cfdrake/go-gdbm"
	pb "github.com/kentaro/grpc-gdbm/gdbm"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var port int
var file string

func init() {
	flag.IntVar(&port, "port", 50051, "port number")
	flag.StringVar(&file, "file", "grpc.gdbm", "gdbm file name")
	flag.Parse()
}

type server struct {
	Db *gdbm.Database
}

func (s *server) Insert(ctx context.Context, in *pb.Request) (*pb.Entry, error) {
	err := s.Db.Insert(in.Key, in.Value)
	return &pb.Entry{Key: in.Key, Value: in.Value}, err
}

func (s *server) Replace(ctx context.Context, in *pb.Request) (*pb.Entry, error) {
	err := s.Db.Replace(in.Key, in.Value)
	return &pb.Entry{Key: in.Key, Value: in.Value}, err
}

func (s *server) Fetch(ctx context.Context, in *pb.Request) (*pb.Entry, error) {
	value, err := s.Db.Fetch(in.Key)
	return &pb.Entry{Key: in.Key, Value: value}, err
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := gdbm.Open(file, "c")
	if err != nil {
		log.Panicf("couldn't open db: %s", err)
	}
	defer db.Close()

	s := grpc.NewServer()
	pb.RegisterGdbmServer(s, &server{Db: db})
	s.Serve(lis)
}
