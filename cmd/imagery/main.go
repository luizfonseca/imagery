package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/luizfonseca/imagery/pkg/handlers"
	"github.com/luizfonseca/imagery/pkg/middleware"
	"go.uber.org/zap"
)

type EnvConfig struct {
	Port string `envconfig:"PORT" default:"4000"`
}

var logger *zap.Logger

func main() {
	var config EnvConfig

	err := envconfig.Process("imagery", &config)
	if err != nil {
		log.Fatalf("Could not process environment configuration %v", err)
	}

	logger, _ = zap.NewProduction()
	client := &http.Client{}
	mux := http.NewServeMux()

	ctxOptions := middleware.NewContextOptions(logger, client)
	mux.HandleFunc("/v1/image", middleware.NewRoute(handlers.ImageHandler, ctxOptions))

	logger.Info(fmt.Sprintf("Server started on port :%s", config.Port))
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), mux)
}
