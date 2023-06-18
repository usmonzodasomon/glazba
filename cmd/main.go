package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/usmonzodasomon/glazba"
	"github.com/usmonzodasomon/glazba/db"
	"github.com/usmonzodasomon/glazba/logger"
	"github.com/usmonzodasomon/glazba/pkg/handler"
	"github.com/usmonzodasomon/glazba/pkg/repository"
	"github.com/usmonzodasomon/glazba/pkg/service"
)

func main() {
	log := logger.GetLogger()
	defer logger.CloseFile()

	if err := initConfig(); err != nil {
		log.Error("Error occured while init viper config: ", err)
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Errorf("error loading env variables: %s", err)
		return
	}

	var dbConfig db.Config
	if err := viper.UnmarshalKey("db", &dbConfig); err != nil {
		log.Error("Error unmarshalling config: ", err)
		return
	}
	dbConfig.Password = os.Getenv("DB_PASSWORD")

	db.StartDbConnection(dbConfig)

	if err := db.Migrate(db.GetDBConn()); err != nil {
		log.Error("Error while migrating tables: ", err)
		return
	}

	repos := repository.NewRepository(db.GetDBConn())
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(glazba.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			log.Error("Error occured while starting server: ", err)
			return
		}
	}()
	log.Info("Server started...")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	if err := db.CloseDbConnection(); err != nil {
		log.Error("Error occured on database connection closing: ", err)
		return
	}

	log.Info("Server closed...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Error("Error server shutting down: ", err)
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
