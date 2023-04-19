package dto

import (
	"context"

	"github.com/skunkie/grpc-gateway-example/internal/pb"
)

type User struct {
	Id    string
	Name  string
	Email string
}

type CreateUserReq struct {
	Name  string
	Email string
}

type UpdateUserReq struct {
	Id    string
	Name  *string
	Email *string
}

type Users interface {
	ListUsers(context.Context, ...string) ([]*User, error)
	CreateUser(context.Context, *CreateUserReq) (*User, error)
	GetUser(context.Context, string) (*User, error)
	UpdateUser(context.Context, *UpdateUserReq) (*User, error)
	DeleteUser(context.Context, string) error
}

func PbCreateUserReqToDtoCreateUserReq(in *pb.CreateUserReq) *CreateUserReq {
	return &CreateUserReq{
		Name:  in.Name,
		Email: in.Email,
	}
}

func PbUpdateUserReqToDtoUpdateUserReq(in *pb.UpdateUserReq) *UpdateUserReq {
	req := &UpdateUserReq{Id: in.Id}
	if in.Name != nil {
		req.Name = in.Name
	}
	if in.Email != nil {
		req.Email = in.Email
	}

	return req
}

func DtoUserToPbUser(in *User) *pb.User {
	return &pb.User{
		Id:    in.Id,
		Name:  in.Name,
		Email: in.Email,
	}
}

func DtoUsersToPbUsers(in []*User) []*pb.User {
	out := []*pb.User{}
	for _, u := range in {
		out = append(out, DtoUserToPbUser(u))
	}

	return out
}
