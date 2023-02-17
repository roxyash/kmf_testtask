package main

import (
	"context"
	"github.com/roxyash/kmf_testtask/pkg/zaplogger"
	"github.com/roxyash/kmf_testtask/proxy_service/config"
	"github.com/roxyash/kmf_testtask/proxy_service/internal/handler"
	"github.com/roxyash/kmf_testtask/proxy_service/internal/repository"
	"github.com/roxyash/kmf_testtask/proxy_service/internal/route"
	"github.com/roxyash/kmf_testtask/proxy_service/internal/service"
	httpserver "github.com/roxyash/kmf_testtask/proxy_service/server/http"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"log"
)

//	@title			PROXY SERVICE
//	@version		1.0
//	@description	this is service, allow user authentication

//	@host		localhost:8000
//	@BasePath	/api/v1

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	log.Println("Initialize config . . .")
	if err := config.InitConfig("proxy_service/config", "local"); err != nil {
		log.Fatalf("Error initialize config: %s", err.Error())
	}

	log.Println("Initialize logger . . .")
	logger := zaplogger.NewZapLogger(viper.GetString("app.logPath"), "")

	logger.Infof("Initialize env files. . .")
	if err := godotenv.Load(viper.GetString("app.envPath")); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	logger.Infof("Initialize packages of service . . .")

	repo := repository.NewRepository()

	services := service.NewService(repo)

	handlers := handler.NewHandler(logger, services)

	routes := route.NewRoute(handlers)

	srv := new(httpserver.Server)

	go func() {
		if err := srv.Run(viper.GetString("app.port"), routes.InitRoutes()); err != nil {
			logger.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

	logger.Infof("%v on http server, port: %v,  Started", viper.GetString("app.serviceName"), viper.GetString("app.port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Infof("%v on http server, port: %v ShuttingDown", viper.GetString("app.serviceName"), viper.GetString("app.port"))
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Fatalf("error occurred on server shutting down :%s", err.Error())
	}
}
