package application

import (
	"fmt"
	"net/http"

	"github.com/gmalheiro/playground/internal/handler"
	"github.com/gmalheiro/playground/internal/loader"
	"github.com/gmalheiro/playground/internal/repository"
	"github.com/gmalheiro/playground/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ConfigServerChi struct {
	ServerAddress  string
	LoaderFilePath string
}

type ServerChi struct {
	serverAddress  string
	loaderFilePath string
}

func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}

	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}
	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

func (a *ServerChi) Run() (err error) {
	ld := loader.NewVehicleJSONFile(a.loaderFilePath)
	db, err := ld.Load()
	if err != nil {
		return
	}
	rp := repository.NewVehicleMap(db)
	sv := service.NewVehicleDefault(rp)
	hd := handler.NewVehicleDefault(sv)
	rt := chi.NewRouter()

	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	rt.Route("/vehicles", func(rt chi.Router) {
		rt.Get("/", hd.GetAll())
	})
	fmt.Printf("Server listening to%s", a.serverAddress)
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
