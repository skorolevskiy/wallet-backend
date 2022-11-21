package main

import (
	"github.com/skorolevskiy/wallet-backend/internal/delivery/http"
	"github.com/skorolevskiy/wallet-backend/internal/server"
	"log"
)

func main() {
	handler := new(http.Handler)
	srv := new(server.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while runing http server: %s", err.Error())
	}
}
