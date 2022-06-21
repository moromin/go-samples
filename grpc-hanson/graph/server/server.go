package main

import (
	"go-samples/grpc-hanson/article/client"
	"go-samples/grpc-hanson/graph"
	"go-samples/grpc-hanson/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// articleClientを作成
	articleClient, err := client.NewClient("localhost:50051")
	if err != nil {
		// articleClient.Close()
		log.Fatalf("failed to create article client: %v\n", err)
	}

	// GraphQLサーバに先程のResolverを実装
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					ArticleClient: articleClient,
				}}))

	// GraphQL playgroundのエンドポイント
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	// 実装したクエリが実行可能なGraphQLのサーバーエンドポイント
	http.Handle("/query", srv)

	// GraphQLサーバーを起動
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
