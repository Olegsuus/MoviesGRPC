package grpc

import (
	"context"
	"github.com/Olegsuus/MovieProto/gen/models/movie"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *MovieServer) Get(ctx context.Context, req *moviepb.GetRequest) (*moviepb.GetResponse, error) {
	const op = "grpc.get"

	logger := slog.With(slog.String("op", op))

	movieModels, err := s.msP.Get(ctx, req.Id)
	if err != nil {
		logger.Error("Ошибка при получении фильма", slog.String("id", req.Id), slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "Ошибка при получении фильма: %v", err)
	}

	grpcMovie, err := s.TranslatorToGrpc(ctx, *movieModels)
	if err != nil {
		logger.Error("Ошибка при переводе из models в grpc", op, err)
		return nil, err
	}

	return grpcMovie, nil
}
