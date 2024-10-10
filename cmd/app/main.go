package main

import (
	"github.com/Olegsuus/MoviesGRPC/internal/app"
	"github.com/Olegsuus/MoviesGRPC/internal/config"
	"github.com/Olegsuus/MoviesGRPC/internal/storage/db"
	storage "github.com/Olegsuus/MoviesGRPC/internal/storage/movie"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mongoStorage, err := db.NewMongoStorage(*cfg)
	if err != nil {
		l.Error("Ошибка подключения к Mongo", slog.Any("error", err))
		os.Exit(1)
	}
	defer mongoStorage.Close()

	movieStorage := storage.RegisterMovieStorage(mongoStorage)

	App := app.New(l, movieStorage, cfg)

	App.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGTERM)

	sign := <-stop
	l.Info("Остановка приложения... Сигнал: %s", sign.String())

	App.Stop()

	l.Info("Приложение остановлено")
}
