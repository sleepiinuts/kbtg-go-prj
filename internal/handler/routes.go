package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sleepiinuts/kbtg-go-prj/internal/ping"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/ping", ping.EchoPing)
	e.POST("/hello", ping.EchoHello)
}
