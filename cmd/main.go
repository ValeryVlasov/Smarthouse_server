package main

import (
	"context"
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/ValeryVlasov/Smarthouse_server/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	handler := new(handler.Handler)

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	srv := new(Smarthouse_server.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("An error occured while starting server: %s", err.Error())
		}
	}()
	logrus.Print("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Server stopped")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("An error occured while server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
