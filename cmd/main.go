package main

import (
	"fmt"
	"point-of-sales-api/internal/config"
	"point-of-sales-api/internal/infra/database"
)

func main() {
	err := config.LoadConfig("./config.yml")
	if err != nil {
		panic(err)
	}

	cfg := config.GetConfig()

	db,err := database.ConnectPostgreSQL(cfg.DB)
	if err != nil {
		panic(err)
	}

	_ = db // Use db as needed

	fmt.Printf("%+v\n", cfg)
}
