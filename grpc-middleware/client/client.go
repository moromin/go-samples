package main

import (
	"context"
	"log"

	"github.com/moromin/go-samples/grpc-middleware/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	md := metadata.New(map[string]string{"authorization": "Bearer hoge"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	c := proto.NewHelloServiceClient(conn)

	res, err := c.Hello(ctx, &proto.HelloRequest{
		Name: "moromin",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response from server: %s\n", res.Message)
}
