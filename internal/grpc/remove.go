package grpc

import (
	"context"
	"github.com/Olegsuus/MovieProto/gen/models/movie"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *MovieServer) Remove(ctx context.Context, req *moviepb.RemoveRequest) (*moviepb.RemoveResponse, error) {
	const op = "grpc.remove"

	logger := s.l.With(slog.String("op", op))

	err := s.msP.Remove(ctx, req.Id)
	if err != nil {
		logger.Error("Ошибка при удалении фильма", "id", req.Id, "error", err)
		return nil, status.Errorf(codes.Internal, "Ошибка при удалении фильма: %v", err)
	}

	return &moviepb.RemoveResponse{
		Status: true,
	}, nil
}
