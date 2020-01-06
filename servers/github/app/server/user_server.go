package server

import (
	"context"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/go-github/v28/github"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (s *userServiceServerImpl) GetUser(ctx context.Context, req *github_pb.GetUserRequest) (*github_pb.User, error) {
	q := req.GetLogin() + " in:login type:user"
	result, _, err := s.gc.Search.Users(ctx, q, nil)
	if err != nil {
		return nil, err
	}
	if result.GetTotal() == 0 {
		return nil, status.Errorf(codes.NotFound, "not found: %q", q)
	}
	return toUserPb(&result.Users[0]), nil
}

func (s *userServiceServerImpl) ListUsers(ctx context.Context, req *github_pb.ListUsersRequest) (*github_pb.ListUsersResponse, error) {
	q := strings.Join(req.GetLogins(), " OR ") + " in:login type:user"
	result, _, err := s.gc.Search.Users(ctx, q, nil)
	if err != nil {
		return nil, err
	}

	userByID := make(map[string]github.User, len(result.Users))
	for _, u := range result.Users {
		userByID[u.GetLogin()] = u
	}

	resp := &github_pb.ListUsersResponse{}
	for _, login := range req.GetLogins() {
		if u, ok := userByID[login]; ok {
			resp.Users = append(resp.Users, toUserPb(&u))
		}
	}

	return resp, nil
}

func toUserPb(u *github.User) *github_pb.User {
	return &github_pb.User{
		Login:                   u.GetLogin(),
		Id:                      u.GetID(),
		NodeId:                  u.GetNodeID(),
		AvatarUrl:               u.GetAvatarURL(),
		HtmlUrl:                 u.GetHTMLURL(),
		GravatarId:              u.GetGravatarID(),
		Name:                    u.GetName(),
		Company:                 u.GetCompany(),
		Blog:                    u.GetBlog(),
		Location:                u.GetLocation(),
		Email:                   u.GetEmail(),
		Hireable:                u.GetHireable(),
		Bio:                     u.GetBio(),
		PublicRepos:             uint32(u.GetPublicRepos()),
		PublicGists:             uint32(u.GetPublicGists()),
		Followers:               uint32(u.GetFollowers()),
		Following:               uint32(u.GetFollowing()),
		CreatedAt:               toTimestamp(u.GetCreatedAt().Time),
		UpdatedAt:               toTimestamp(u.GetUpdatedAt().Time),
		SuspendedAt:             toTimestamp(u.GetSuspendedAt().Time),
		Type:                    u.GetType(),
		SiteAdmin:               u.GetSiteAdmin(),
		TotalPrivateRepos:       uint32(u.GetTotalPrivateRepos()),
		OwnedPrivateRepos:       uint32(u.GetOwnedPrivateRepos()),
		PrivateGists:            uint32(u.GetPrivateGists()),
		DiskUsage:               uint32(u.GetDiskUsage()),
		Collaborators:           uint32(u.GetCollaborators()),
		TwoFactorAuthentication: u.GetTwoFactorAuthentication(),
		Plan:                    toPlanPb(u.GetPlan()),
		LdapDn:                  u.GetLdapDn(),
	}
}

func toTimestamp(t time.Time) *timestamp.Timestamp {
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		panic(err) // maybe unreachable
	}
	return ts
}

func toPlanPb(p *github.Plan) *github_pb.Plan {
	if p == nil {
		return nil
	}
	return &github_pb.Plan{
		Name:          p.GetName(),
		Space:         uint32(p.GetSpace()),
		Collaborators: uint32(p.GetCollaborators()),
		PrivateRepos:  uint32(p.GetPrivateRepos()),
		// TODO: not supported in go-github
		// FilledSeats:   p.GetFilledSeats(),
		// Seats:         p.GetSeats(),
	}
}
