package main

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
	"github.com/witalok2/api-rest/adapters/primary/api"
	"github.com/witalok2/api-rest/adapters/secondary/postgres"
	"github.com/witalok2/api-rest/config"
	"github.com/witalok2/api-rest/internal/domain/user"
	"github.com/witalok2/api-rest/internal/entity"
)

func init() {
	logger.New().WithContext(context.WithValue(context.Background(), entity.SERVICE, "api-rest"))
}

func main() {
	env, err := config.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}

	repository := postgres.New(env.Postgres)
	defer repository.Close()

	startedTime := time.Now()
	logger.Debugf("Starting execution: %v", startedTime)

	userService := user.NewService(repository)

	api.New()
	api.ProvideEchoInstance(func(e *echo.Echo) { api.AddRoutes(e, userService) })
	api.Run()

}
