package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log/slog"
	"time"
)

func LoggingInterceptor(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()
		logger.Info("Начало gRPC вызова",
			slog.String("method", info.FullMethod),
			slog.Any("request", req),
		)

		resp, err := handler(ctx, req)

		if err != nil {
			logger.Error("Ошибка в gRPC вызове",
				slog.String("method", info.FullMethod),
				slog.Any("error", err),
			)
		} else {
			logger.Info("Успешный gRPC вызов",
				slog.String("method", info.FullMethod),
				slog.Any("response", resp),
				slog.Duration("duration", time.Since(start)),
			)
		}

		return resp, err
	}
}
