package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/kikytokamuro/grpc_example/space"
)

var (
	addr = flag.String("addr", "localhost:5030", "The address to connect to")
	str  = flag.String("str", "Test String", "Spacify string")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Not connected: %v", err)
	}
	defer conn.Close()

	client := pb.NewSpacerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Do(ctx, &pb.Request{Message: *str})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Result: %v", resp.GetMessage())
}
