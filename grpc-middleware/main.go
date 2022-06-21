package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/moromin/go-samples/grpc-middleware/grpc"
)

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		errCh <- grpc.RunServer(ctx, 5000)
	}()

	select {
	case err := <-errCh:
		log.Println(err)
		return 1
	case <-ctx.Done():
		log.Println("shutting down...")
		return 0
	}
}
