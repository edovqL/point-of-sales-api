package main

import (
	"log/slog"
	"os"
	"point-of-sales-api/internal/config"
	"point-of-sales-api/internal/server"
)

func main() {
	err := config.LoadConfig("./config.yml")
	if err != nil {
		panic(err)
	}

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: false,
	})

	log := slog.New(logHandler)
	slog.SetDefault(log)

	if err := server.Start(); err != nil {
		panic(err)
	}

}
