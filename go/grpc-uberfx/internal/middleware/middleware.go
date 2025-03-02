package middleware

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// LoggingUnaryInterceptor untuk logging request dan response di gRPC
func LoggingUnaryInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()

		// Log request masuk
		logger.Info("Incoming gRPC request",
			zap.String("method", info.FullMethod),
			zap.Any("request", req),
		)

		// Panggil handler berikutnya
		resp, err := handler(ctx, req)

		// Log response keluar
		logger.Info("Outgoing gRPC response",
			zap.String("method", info.FullMethod),
			zap.Any("response", resp),
			zap.Duration("duration", time.Since(start)),
			zap.Error(err),
			zap.String("status", status.Code(err).String()),
		)

		return resp, err
	}
}
