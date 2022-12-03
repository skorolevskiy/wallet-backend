package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/skorolevskiy/wallet-backend/internal/config"
	delivery "github.com/skorolevskiy/wallet-backend/internal/delivery/rest"
	"github.com/skorolevskiy/wallet-backend/internal/repository"
	"github.com/skorolevskiy/wallet-backend/internal/server"
	"github.com/skorolevskiy/wallet-backend/internal/service"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variable: %s", err.Error())
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
		logrus.Fatalf("failed to install DB: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := delivery.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while runing http server: %s", err.Error())
	}
}
