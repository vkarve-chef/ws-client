package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "ws-client/io.chef"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:56001", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChefInfraClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetCookbookVersion(ctx, &pb.Cookbook{Name: "zfs"})
	if err != nil {
		log.Fatalf("could not get cookbook version: %v", err)
	}
	log.Printf("Version found: %d.%d.%d", r.GetMajor(), r.GetMinor(), r.GetPatch())
}
