package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	github_pb "github.com/izumin5210-sandbox/rejoiner-apollo-sample/api/go/github"
)

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *userServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	github_pb.RegisterUserServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *userServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return nil
}
