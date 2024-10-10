package app

import (
	"fmt"
	"github.com/Olegsuus/MoviesGRPC/internal/config"
	grpc2 "github.com/Olegsuus/MoviesGRPC/internal/grpc"
	"github.com/Olegsuus/MoviesGRPC/internal/services"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	cfg        *config.Config
}

func New(log *slog.Logger, movieService services.MovieService, cfg *config.Config) *App {
	gRPCServer := grpc.NewServer()

	grpc2.Register(gRPCServer, movieService) //to

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		cfg:        cfg,
	}
}

func (a *App) Run() error {
	const op = "app.Run"

	log := a.log.With(slog.String("op", op))

	port := a.cfg.App.GRPC.Port

	log.Info("starting gRPC server", slog.Int("port", port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gPRC server is running", slog.String("addr", l.Addr().String()))

	if err = a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	const op = "app.Stop"
	a.log.With(slog.String("op", op)).Info("stopping gRPC server")
	a.gRPCServer.GracefulStop()
}
