package main

import (
	"flag"
	"log"
	"net"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"github.com/bur8787/go-gke-example/src/product"
)


var (
	addr = flag.String("addr", ":8080", "Network host:port to listen on for gRPC connections.")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	product.RegisterProductServer(s, &product.Server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

