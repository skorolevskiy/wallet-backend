package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/skorolevskiy/wallet-backend/internal/config"
	delivery "github.com/skorolevskiy/wallet-backend/internal/delivery/http"
	"github.com/skorolevskiy/wallet-backend/internal/repository"
	"github.com/skorolevskiy/wallet-backend/internal/server"
	"github.com/skorolevskiy/wallet-backend/internal/service"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error initialization configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variable: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.ports"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASS"),
	})
	if err != nil {
		log.Fatalf("failed to install DB: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := delivery.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run(viper.GetString("8000"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while runing http server: %s", err.Error())
	}
}
