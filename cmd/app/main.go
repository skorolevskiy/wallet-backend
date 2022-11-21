package main

import (
	"github.com/skorolevskiy/wallet-backend/internal/config"
	delivery "github.com/skorolevskiy/wallet-backend/internal/delivery/http"
	"github.com/skorolevskiy/wallet-backend/internal/repository"
	"github.com/skorolevskiy/wallet-backend/internal/server"
	"github.com/skorolevskiy/wallet-backend/internal/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error initialization configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := delivery.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run(viper.GetString("8000"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while runing http server: %s", err.Error())
	}
}
