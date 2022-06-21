package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/moromin/go-samples/grpc-middleware/proto"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

func RunServer(ctx context.Context, port int) error {
	// auth
	// s := grpc.NewServer(
	// 	grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(auth.Validate)),
	// )

	// logger(zap)
	zapLogger, err := New()
	if err != nil {
		return err
	}
	zapOpts := grpc_zap.WithLevels(
		func(c codes.Code) zapcore.Level {
			var l zapcore.Level
			switch c {
			case codes.OK:
				l = zapcore.InfoLevel
			case codes.Internal:
				l = zapcore.ErrorLevel
			default:
				l = zapcore.DebugLevel
			}
			return l
		},
	)

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			ctxtags.UnaryServerInterceptor(ctxtags.WithFieldExtractor(ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(zapLogger, zapOpts),
		),
	)

	proto.RegisterStudentServer(s, &server{})
	reflection.Register(s)

	log.Printf("server listening port: %d", port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	return s.Serve(lis)
}
