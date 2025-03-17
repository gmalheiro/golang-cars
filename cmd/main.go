package main

import (
	"fmt"

	"github.com/gmalheiro/playground/internal/application"
)

func main() {
	cfg := &application.ConfigServerChi{
		ServerAddress:  ":8080",
		LoaderFilePath: "/Users/gsmalheiro/studies/go/playground/docs/db/vehicles_100.json",
	}
	app := application.NewServerChi(cfg)
	if err := app.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
