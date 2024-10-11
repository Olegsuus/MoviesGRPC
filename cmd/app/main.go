package main

import (
	"github.com/Olegsuus/MoviesGRPC/internal/app"
	"github.com/Olegsuus/MoviesGRPC/internal/config"
	"github.com/Olegsuus/MoviesGRPC/internal/services"
	"github.com/Olegsuus/MoviesGRPC/internal/storage/db"
	storage "github.com/Olegsuus/MoviesGRPC/internal/storage/movie"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mongoStorage, err := db.NewMongoStorage(*cfg)
	if err != nil {
		logger.Error("Ошибка подключения к Mongo", slog.Any("error", err))
		os.Exit(1)
	}
	defer mongoStorage.Close()

	movieStorage := storage.RegisterMovieStorage(mongoStorage)
	movieService := services.RegisterMovieService(logger, movieStorage)

	App := app.New(logger, movieService, cfg)

	go func() {
		App.MustRun()
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	logger.Info("Остановка приложения... Сигнал: " + sign.String())

	App.Stop()

	logger.Info("Приложение остановлено")
}
