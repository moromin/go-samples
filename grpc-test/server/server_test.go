package main

import (
	"context"
	"go-samples/grpc-test/proto"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
}

func bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return lis.Dial()
}

// func TestSayHello(t *testing.T) {
// 	ctx := context.Background()
// 	opts := []grpc.DialOption{
// 		grpc.WithContextDialer(bufDialer),
// 		grpc.WithInsecure(),
// 	}

// 	conn, err := grpc.DialContext(ctx, "bufnet", opts...)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer conn.Close()

// 	client := proto.NewGreeterClient(conn)
// 	resp, err := client.SayHello(ctx, &proto.HelloRequest{Name: "test"})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if resp.GetMessage() != "Hello test" {
// 		t.Fatal("hello reply must be 'Hello test'")
// 	}
// }

func Test_server_SayHello(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *proto.HelloRequest
	}

	ctx := context.Background()

	tests := []struct {
		name string
		args args
		want *proto.HelloReply
	}{
		{
			name: "moromin",
			args: args{ctx, &proto.HelloRequest{Name: "moromin"}},
			want: &proto.HelloReply{Message: "Hello moromin"},
		},
		{
			name: "takuro",
			args: args{ctx, &proto.HelloRequest{Name: "takuro"}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{}
			got, err := s.SayHello(tt.args.ctx, tt.args.in)
			if err != nil {
				st, ok := status.FromError(err)
				if ok && st.Code() != codes.InvalidArgument {
					t.Errorf("got code: %v", st.Code())
				}
				t.Errorf("got error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
