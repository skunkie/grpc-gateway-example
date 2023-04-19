package dto

import (
	"context"

	"github.com/skunkie/grpc-gateway-example/internal/pb"
)

type Token struct {
	TokenString string
}

type CreateTokenReq struct {
	Username string
	Password string
}

type Auth interface {
	CreateToken(context.Context, *CreateTokenReq) (*Token, error)
}

func PbCreateTokenReqToDtoCreateTokenReq(in *pb.CreateTokenReq) *CreateTokenReq {
	return &CreateTokenReq{
		Username: in.Username,
		Password: in.Password,
	}
}

func DtoTokenToPbToken(in *Token) *pb.Token {
	return &pb.Token{
		TokenString: in.TokenString,
	}
}
