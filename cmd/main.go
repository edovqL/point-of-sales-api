package main

import (
	"fmt"
	"point-of-sales-api/internal/config"
)

func main() {
	err := config.LoadConfig("./config.yml")
	if err != nil {
		panic(err)
	}

	cfg := config.GetConfig()

	fmt.Printf("%+v\n", cfg)
}
