package controller

import (
	"aiplus/internal/app/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Controller struct {
	s service.Service
}

func NewController(s service.Service) Controller {
	return Controller{
		s: s,
	}
}

func (c Controller) Build() *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.RequestID(),
		middleware.Recover(),
		middleware.CORS(),
	)

	e.Any("/swagger/*", echoSwagger.WrapHandler)

	{
		staff := e.Group("/staff", middleware.Logger())
		staff.POST("", c.createStaff)
	}

	return e
}
