package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/skunkie/grpc-gateway-example/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const authFullMethod = "/pb.Api/CreateToken"

var (
	signingKey   = []byte("secret")
	tokenInfoKey struct{}
)

type Auth struct{}

func New() *Auth {
	return &Auth{}
}

func (a *Auth) CreateToken(ctx context.Context, in *dto.CreateTokenReq) (*dto.Token, error) {
	claims := &jwt.RegisteredClaims{
		Issuer:    "example.com",
		Subject:   in.Username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return nil, err
	}

	return &dto.Token{TokenString: tokenString}, nil
}

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func userClaimFromToken(token *jwt.Token) string {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		sub, err := claims.GetSubject()
		if err != nil {
			return ""
		}
		return sub
	}

	return ""
}

// AuthFunc is used by a middleware to authenticate requests
func AuthFunc(ctx context.Context) (context.Context, error) {
	tokenString, err := grpcauth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	tokenInfo, err := parseToken(tokenString)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	ctx = logging.InjectFields(ctx, logging.Fields{"auth.sub", userClaimFromToken(tokenInfo)})

	return context.WithValue(ctx, tokenInfoKey, tokenInfo), nil
}

func TokenOnly(_ context.Context, c interceptors.CallMeta) bool {
	return c.FullMethod() == authFullMethod
}

func AllButToken(_ context.Context, c interceptors.CallMeta) bool {
	return c.FullMethod() != authFullMethod
}
