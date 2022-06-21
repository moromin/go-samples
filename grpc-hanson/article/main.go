// 動作確認用
package main

import (
	"context"
	"fmt"
	"go-samples/grpc-hanson/article/client"
	"go-samples/grpc-hanson/article/pb"
	"io"
	"log"
)

func main() {
	c, _ := client.NewClient("localhost:50051")

	create(c)
	read(c)
	update(c)
	delete(c)
	list(c)
}

func create(c *client.Client) {
	input := &pb.ArticleInput{
		Author:  "gopher-san",
		Title:   "gRPC-kun",
		Content: "gRPC is so great!",
	}

	res, err := c.Service.CreateArticle(context.Background(), &pb.CreateArticleRequest{ArticleInput: input})
	if err != nil {
		log.Fatal("failed to CreateArticle:", err)
	}
	fmt.Printf("CreateArticle Response: %v\n", res)
}

func read(c *client.Client) {
	var id int64 = 7
	res, err := c.Service.ReadArticle(context.Background(), &pb.ReadArticleRequest{Id: id})
	if err != nil {
		log.Fatal("failed to ReadArticle:", err)
	}
	fmt.Printf("ReadArticle Response: %v\n", res)
}

func update(c *client.Client) {
	var id int64 = 7
	input := &pb.ArticleInput{
		Author:  "GraphQL master",
		Title:   "GraphQL",
		Content: "GraphQL is very smart!",
	}
	res, err := c.Service.UpdateArticle(context.Background(), &pb.UpdateArticleRequest{ArticleInput: input, Id: id})
	if err != nil {
		log.Fatal("failed to UpdateArticle:", err)
	}
	fmt.Printf("UpdateArticle Response: %v\n", res)
}

func delete(c *client.Client) {
	var id int64 = 6
	res, err := c.Service.DeleteArticle(context.Background(), &pb.DeleteArticleRequest{Id: id})
	if err != nil {
		log.Fatal("failed to DeleteArticle:", err)
	}
	fmt.Printf("DeleteArticle Response: %v\n", res)
}

func list(c *client.Client) {
	stream, err := c.Service.ListArticle(context.Background(), &pb.ListArticleRequest{})
	if err != nil {
		log.Fatal("failed to ListArticle:", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("failed to server streaming:", err)
		}
		fmt.Println(res)
	}
}
