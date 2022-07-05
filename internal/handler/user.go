package domain

import (
	"net/http"

	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
	"github.com/witalok2/api-rest/internal/domain/user"
	"github.com/witalok2/api-rest/internal/entity"
)

func CreateUser(s user.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		body := new(entity.User)
		if err := c.Bind(body); err != nil {
			logger.New().WithError(err).Error("failed to parse request body")
			return c.JSON(http.StatusInternalServerError, err)
		}

		if err := body.Validate(); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		err := s.CreateUser(ctx, *body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, nil)
	}
}

func ListUser(s user.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		users, err := s.ListUser(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, users)
	}
}
