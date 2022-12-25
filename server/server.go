package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"regexp"

	pb "github.com/kikytokamuro/grpc_example/space"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 5030, "The server port")
)

type server struct {
	pb.UnimplementedSpacerServer
}

func (s *server) Do(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	whitespaces := regexp.MustCompile(`\s+`)
	replaced := whitespaces.ReplaceAllString(in.GetMessage(), "_")

	log.Printf("Request: %v\n Response: %v", in.GetMessage(), replaced)

	return &pb.Response{Message: replaced}, nil
}

func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterSpacerServer(srv, &server{})

	log.Printf("Server listening at %v", listen.Addr())

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
