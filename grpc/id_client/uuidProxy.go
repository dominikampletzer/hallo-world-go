package main

import (
	"context"
	"github.com/dominikampletzer/hello/grpc/helloworld"
	"google.golang.org/grpc"
	"log"
	"time"
)

func GetUuid() int64 {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	receivedUuid, err := c.GetUUID(ctx, &helloworld.UUIDRequest{})
	if err != nil {
		log.Fatalf("could not receive uuid: %v", err)
	}
	log.Printf("UUID: %v", receivedUuid.Uuid)
	return receivedUuid.Uuid
}
