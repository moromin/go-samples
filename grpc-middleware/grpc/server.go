package grpc

import (
	"context"

	"github.com/moromin/go-samples/grpc-middleware/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ proto.StudentServer = (*server)(nil)

type server struct{}

func (s *server) Get(ctx context.Context, req *proto.StudentRequest) (*proto.StudentResponse, error) {
	if req.Id == 1 {
		res := &proto.StudentResponse{
			Id:   1,
			Name: "Taro",
			Age:  11,
			School: &proto.School{
				Id:    1,
				Name:  "ABC school",
				Grade: "5th",
			},
		}
		// ログ出力するフィールドを定義
		// log := map[string]interface{}{ // --- ②
		// 	"name":         res.Name,
		// 	"age":          res.Age,
		// 	"school_name":  res.School.Name,
		// 	"school_grade": res.School.Grade,
		// }
		// grpc_ctxtags.Extract(ctx).Set("data", log)
		return res, nil
	} else {
		// grpc_ctxtags.Extract(ctx).Set("request_id", req.Id)             // --- ①
		return nil, status.Errorf(codes.Internal, "No students found.") // --- ③
	}
}
