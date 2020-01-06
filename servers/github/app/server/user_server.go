package server

import (
	"context"

	"github.com/google/go-github/v28/github"
	"github.com/izumin5210/grapi/pkg/grapiserver"

	github_pb "github.com/izumin5210-sandbox/rejoiner-apollo-sample/api/go/github"
)

// UserServiceServer is a composite interface of api_pb.UserServiceServer and grapiserver.Server.
type UserServiceServer interface {
	grapiserver.Server
	github_pb.UserServiceServer
}

// NewUserServiceServer creates a new UserServiceServer instance.
func NewUserServiceServer(
	gc *github.Client,
) UserServiceServer {
	return &userServiceServerImpl{
		gc: gc,
	}
}

type userServiceServerImpl struct {
	gc *github.Client
}

func (s *userServiceServerImpl) GetUser(context.Context, *github_pb.GetUserRequest) (*github_pb.User, error) {
	return nil, nil
}

func (s *userServiceServerImpl) ListUser(context.Context, *github_pb.ListUserRequest) (*github_pb.ListUserResponse, error) {
	return nil, nil
}
