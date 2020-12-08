package main

import (
	"fmt"
	"log"
	"net"
	"net/url"

	"solution/infrastructure"
	"solution/proto/serviceBook"
	handler "solution/serviceBook"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :8082...")

	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("Unable to listen on port :8082: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	srv := &handler.BookServiceServer{}

	bookpb.RegisterBookServiceServer(s, srv)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func init() {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	env := infrastructure.Environment{}
	env.SetEnvironment()
	env.LoadConfig()
	env.InitMongoDB()
}
