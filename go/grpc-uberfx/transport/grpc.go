package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"

	"github.com/Arif9878/design-chat-system/go/grpc-uberfx/proto"
	"github.com/Arif9878/design-chat-system/go/grpc-uberfx/service"
)

// Implementasi gRPC server
type grpcServer struct {
	proto.UnimplementedHelloServiceServer
	sayHello kitgrpc.Handler
}

// SayHello gRPC handler
func (s *grpcServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	_, resp, err := s.sayHello.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.HelloResponse), nil
}

// NewGRPCServer membuat instance gRPC server
func NewGRPCServer(svc service.Service, logger *zap.Logger) proto.HelloServiceServer {
	return &grpcServer{
		sayHello: kitgrpc.NewServer(
			makeSayHelloEndpoint(svc, logger),
			decodeHelloRequest,
			encodeHelloResponse,
		),
	}
}

// Endpoint untuk gRPC
func makeSayHelloEndpoint(svc service.Service, logger *zap.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.HelloRequest)
		msg, err := svc.SayHello(ctx, req.Name)
		if err != nil {
			return nil, err
		}
		return &proto.HelloResponse{Message: msg}, nil
	}
}

// Decode request gRPC
func decodeHelloRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req.(*proto.HelloRequest), nil
}

// Encode response gRPC
func encodeHelloResponse(_ context.Context, resp interface{}) (interface{}, error) {
	return resp.(*proto.HelloResponse), nil
}
