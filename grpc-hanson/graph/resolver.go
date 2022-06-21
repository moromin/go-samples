package graph

import (
	"go-samples/grpc-hanson/article/client"
)

type Resolver struct {
	ArticleClient *client.Client
}
