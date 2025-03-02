package grpc

import (
	"context"
	"net"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/Arif9878/design-chat-system/go/grpc-uberfx/cmd/initialize"
	"github.com/Arif9878/design-chat-system/go/grpc-uberfx/internal/config"
	"github.com/Arif9878/design-chat-system/go/grpc-uberfx/internal/logging"
	"github.com/Arif9878/design-chat-system/go/grpc-uberfx/internal/middleware"
	"github.com/Arif9878/design-chat-system/go/grpc-uberfx/proto"
	"github.com/Arif9878/design-chat-system/go/grpc-uberfx/service"
	"github.com/Arif9878/design-chat-system/go/grpc-uberfx/transport"
)

// gRPC Server Lifecycle
func NewGRPCServer(lc fx.Lifecycle, svc service.Service, logger *zap.Logger) *grpc.Server {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(middleware.LoggingUnaryInterceptor(logger)),
	)
	proto.RegisterHelloServiceServer(server, transport.NewGRPCServer(svc, logger))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lis, err := net.Listen("tcp", ":50051")
			if err != nil {
				return err
			}
			logger.Info("gRPC server started on :50051")
			go server.Serve(lis)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down gRPC server")
			server.GracefulStop()
			return nil
		},
	})

	return server
}

func Exec(ctx context.Context, cfg *config.DBConfig) func(clx *cli.Context) error {
	return func(clx *cli.Context) error {
		app := fx.New(
			fx.Provide(func() *config.DBConfig { return cfg }), // Inject config
			logging.Module,
			initialize.DBModule,
			fx.Provide(
				service.NewService,
				NewGRPCServer,
			),
			fx.Invoke(func(*grpc.Server) {}),
		)

		err := app.Start(ctx) // start app
		if err != nil {
			return err
		}

		<-ctx.Done() // wait signal to stop
		err = app.Stop(ctx)
		return err
	}
}
