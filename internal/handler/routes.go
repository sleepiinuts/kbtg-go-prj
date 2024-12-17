package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sleepiinuts/kbtg-go-prj/internal/ping"
)

type route struct {
	Group      string
	Path       string
	Method     string
	Endpoint   echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
}

func InitRoutes(e *echo.Echo) {
	routes := []route{
		{
			Path:       "/ping",
			Method:     http.MethodGet,
			Endpoint:   ping.EchoPing,
			Middleware: []echo.MiddlewareFunc{MovePermanant},
		},
		{
			Path:     "/hello",
			Method:   http.MethodPost,
			Endpoint: ping.EchoHello,
		},
	}

	for _, r := range routes {
		e.Group(r.Group).Add(r.Method, r.Path, r.Endpoint, r.Middleware...)
	}
	// e.GET("/ping", ping.EchoPing)
	// e.POST("/hello", ping.EchoHello)
}
