package chat

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, req *MessageRequest) (*MessageResponse, error) {
	log.Printf("Receive message body from client: %s", req.Body)
	return &MessageResponse{Body: "Hello From the Server!"}, nil
}
