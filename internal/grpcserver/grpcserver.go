package grpcserver

import (
	"net"

	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"github.com/skunkie/grpc-gateway-example/internal/auth"
	"github.com/skunkie/grpc-gateway-example/internal/pb"
	"github.com/skunkie/grpc-gateway-example/pkg/interceptors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	controller pb.ApiServer
	*grpc.Server
	logger *zap.Logger
}

func New(controller pb.ApiServer, logger *zap.Logger) *Server {
	authOpts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall, logging.PayloadReceived, logging.PayloadSent),
	}

	srv := &Server{
		controller: controller,
		Server: grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				selector.UnaryServerInterceptor(grpcauth.UnaryServerInterceptor(auth.AuthFunc), selector.MatchFunc(auth.AllButToken)),
				selector.UnaryServerInterceptor(logging.UnaryServerInterceptor(interceptors.InterceptorLogger(logger), authOpts...), selector.MatchFunc(auth.TokenOnly)),
				selector.UnaryServerInterceptor(logging.UnaryServerInterceptor(interceptors.InterceptorLogger(logger), opts...), selector.MatchFunc(auth.AllButToken)),
				validator.UnaryServerInterceptor(),
				recovery.UnaryServerInterceptor(),
			),
			grpc.ChainStreamInterceptor(
				selector.StreamServerInterceptor(grpcauth.StreamServerInterceptor(auth.AuthFunc), selector.MatchFunc(auth.AllButToken)),
				selector.StreamServerInterceptor(logging.StreamServerInterceptor(interceptors.InterceptorLogger(logger), authOpts...), selector.MatchFunc(auth.TokenOnly)),
				selector.StreamServerInterceptor(logging.StreamServerInterceptor(interceptors.InterceptorLogger(logger), opts...), selector.MatchFunc(auth.AllButToken)),
				validator.StreamServerInterceptor(),
				recovery.StreamServerInterceptor(),
			),
		),
		logger: logger,
	}
	reflection.Register(srv)

	return srv
}

func (s *Server) Run(lis net.Listener) error {
	pb.RegisterApiServer(s, s.controller)
	s.logger.Sugar().Infof("gRPC server listening at %v", lis.Addr())
	return s.Serve(lis)
}
