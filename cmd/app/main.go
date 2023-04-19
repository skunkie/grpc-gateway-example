package main

import (
	"context"
	"database/sql"
	"flag"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/proullon/ramsql/driver"
	"github.com/skunkie/grpc-gateway-example/internal/auth"
	"github.com/skunkie/grpc-gateway-example/internal/controller"
	"github.com/skunkie/grpc-gateway-example/internal/grpcserver"
	"github.com/skunkie/grpc-gateway-example/internal/httpserver"
	"github.com/skunkie/grpc-gateway-example/internal/middleware"
	"github.com/skunkie/grpc-gateway-example/internal/pb"
	"github.com/skunkie/grpc-gateway-example/internal/repository/users"
	"github.com/skunkie/grpc-gateway-example/internal/swagger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcAddr = flag.String("g", ":0", "gRPC address")
	httpAddr = flag.String("h", ":8080", "HTTP address")
)

func main() {
	flag.Parse()
	cfg := httpserver.Config{Addr: *httpAddr}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	logger := zap.NewExample()
	defer logger.Sync()
	log := logger.Sugar()

	lis, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	mux := runtime.NewServeMux(
		runtime.WithForwardResponseOption(middleware.HTTPStatusCode),
		runtime.WithForwardResponseOption(middleware.HTTPResponseModifier),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = pb.RegisterApiHandlerFromEndpoint(ctx, mux, lis.Addr().String(), opts)
	if err != nil {
		log.Fatal(err)
	}
	httpServer := httpserver.New(cfg, mux, swagger.GetDocs(), logger)

	db, err := sql.Open("ramsql", "db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE users (id TEXT PRIMARY KEY, name TEXT, email TEXT);`)
	if err != nil {
		log.Fatalf("sql.Exec: Error: %s\n", err)
	}

	c := controller.New(auth.New(), users.New(db))
	grpcServer := grpcserver.New(c, logger)
	go func() {
		if err := grpcServer.Run(lis); err != nil {
			panic(err)
		}
	}()

	// Shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Info("shutting down...")
	toCtx, toCancel := context.WithTimeout(context.Background(), time.Second*15)
	defer toCancel()

	if err = httpServer.Shutdown(toCtx); err != nil {
		log.Fatalf("failed to shut down http server: %s", err)
	}
	grpcServer.GracefulStop()
}
