package api

import (
	"github.com/labstack/echo/v4"
	"github.com/witalok2/api-rest/internal/domain/user"
	handler "github.com/witalok2/api-rest/internal/handler"
)

func AddRoutes(e *echo.Echo, s user.Service) {
	v1 := e.Group("/v1")

	v1.POST("/user", handler.CreateUser(s))
	v1.GET("/user", handler.ListUser(s))
}
