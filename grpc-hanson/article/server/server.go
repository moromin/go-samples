package main

import (
	"go-samples/grpc-hanson/article/pb"
	"go-samples/grpc-hanson/article/repository"
	"go-samples/grpc-hanson/article/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("failed to listen:", err)
	}
	defer lis.Close()

	repository, err := repository.NewSqliteRepo()
	if err != nil {
		log.Fatal("failed to create sqlite repository", err)
	}

	service := service.NewService(repository)

	server := grpc.NewServer()
	pb.RegisterArticleServiceServer(server, service)

	log.Println("Listening on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}
