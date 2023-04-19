package controller

import (
	"context"
	"net/http"

	"github.com/skunkie/grpc-gateway-example/internal/dto"
	mw "github.com/skunkie/grpc-gateway-example/internal/middleware"
	"github.com/skunkie/grpc-gateway-example/internal/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type controller struct {
	pb.UnimplementedApiServer
	auth  dto.Auth
	users dto.Users
}

func New(auth dto.Auth, users dto.Users) *controller {
	return &controller{
		auth:  auth,
		users: users,
	}
}

func (c *controller) CreateToken(ctx context.Context, in *pb.CreateTokenReq) (*pb.Token, error) {
	t, err := c.auth.CreateToken(ctx, dto.PbCreateTokenReqToDtoCreateTokenReq(in))
	if err != nil {
		return nil, err
	}

	return dto.DtoTokenToPbToken(t), nil
}

func (c *controller) ListUsers(ctx context.Context, in *pb.Ids) (*pb.Users, error) {
	u, err := c.users.ListUsers(ctx, in.Id...)
	if err != nil {
		return nil, err
	}

	return &pb.Users{Users: dto.DtoUsersToPbUsers(u)}, nil
}

func (c *controller) CreateUser(ctx context.Context, in *pb.CreateUserReq) (*pb.User, error) {
	u, err := c.users.CreateUser(ctx, dto.PbCreateUserReqToDtoCreateUserReq(in))
	if err != nil {
		return nil, err
	}

	return dto.DtoUserToPbUser(u), nil
}

func (c *controller) GetUser(ctx context.Context, in *pb.Id) (*pb.User, error) {
	u, err := c.users.GetUser(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return dto.DtoUserToPbUser(u), nil
}

func (c *controller) UpdateUser(ctx context.Context, in *pb.UpdateUserReq) (*pb.User, error) {
	u, err := c.users.UpdateUser(ctx, dto.PbUpdateUserReqToDtoUpdateUserReq(in))
	if err != nil {
		return nil, err
	}

	mw.SetHTTPStatusCode(ctx, http.StatusOK)
	return dto.DtoUserToPbUser(u), nil
}

func (c *controller) DeleteUser(ctx context.Context, in *pb.Id) (*emptypb.Empty, error) {
	err := c.users.DeleteUser(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
