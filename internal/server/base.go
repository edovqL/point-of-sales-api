package server

import (
	"log/slog"
	"net/http"
	"point-of-sales-api/internal/config"
	"point-of-sales-api/internal/infra/database"

	"github.com/go-chi/chi"
)

func Start() error {
	cfg := config.GetConfig()

	db, err := database.ConnectPostgreSQL(cfg.DB)
	if err != nil {
		panic(err)
	}

	_ = db // Use db as needed

	// fmt.Printf("%+v\n", cfg) Optional

	router := chi.NewRouter()

	slog.Info("Server "+cfg.App.Name+" starting", slog.String("port", cfg.App.Port))

	http.ListenAndServe(":"+cfg.App.Port, router)

	return nil

}
