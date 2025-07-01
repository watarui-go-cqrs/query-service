package interceptor

import (
	"context"
	"fmt"
	"regexp"

	"github.com/watarui-go-cqrs/pb/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func isUUID(id string) bool {
	var uuidPattern = regexp.MustCompile(
		`^([0-9a-fA-F]{8})-([0-9a-fA-F]{4})-([0-9a-fA-F]{4})-([0-9a-fA-F]{4})-([0-9a-fA-F]{12})$`,
	)
	return uuidPattern.MatchString(id)
}

func UUIDValidationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var id string
	switch info.FullMethod {
	case "/proto.CategoryQuery/ById":
		param, _ := req.(*pb.CategoryParam)
		id = param.GetId()
	case "/proto.ProductQuery/ById":
		param, _ := req.(*pb.ProductParam)
		id = param.GetId()
	}
	if id != "" && !isUUID(id) {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Invalid UUID format: %s", id))
	}
	return handler(ctx, req)
}
