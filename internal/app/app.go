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

func New(logger *slog.Logger, movieService *services.MovieService, cfg *config.Config) *App {
	logInterceptor := grpc2.LoggingInterceptor(logger)
	gRPCServer := grpc.NewServer(
		grpc.UnaryInterceptor(logInterceptor),
	)

	grpc2.Register(gRPCServer, movieService)

	return &App{
		log:        logger,
		gRPCServer: gRPCServer,
		cfg:        cfg,
	}
}

func (a *App) Run() error {
	const op = "app.Run"

	log := a.log.With(slog.String("op", op))

	port := a.cfg.App.GRPC.Port

	log.Info("Запуск gRPC сервера", slog.Int("port", port))

	addr := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("%s: ошибка при прослушивании адреса %s: %w", op, addr, err)
	}

	log.Info("gRPC сервер запущен", slog.String("address", lis.Addr().String()))

	if err = a.gRPCServer.Serve(lis); err != nil {
		return fmt.Errorf("%s: ошибка при запуске gRPC сервера: %w", op, err)
	}

	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		a.log.Error("Не удалось запустить сервер", slog.Any("error", err))
		panic(err)
	}
}

func (a *App) Stop() {
	const op = "app.Stop"
	a.log.With(slog.String("op", op)).Info("Остановка gRPC сервера")
	a.gRPCServer.GracefulStop()
}
