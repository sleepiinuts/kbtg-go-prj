package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sleepiinuts/kbtg-go-prj/internal/app"
	"github.com/sleepiinuts/kbtg-go-prj/internal/ping"
	"github.com/sleepiinuts/kbtg-go-prj/internal/school"
	"github.com/sleepiinuts/kbtg-go-prj/internal/student"
)

type route struct {
	Group      string
	Path       string
	Method     string
	Endpoint   echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
}

func InitRoutes(e *echo.Echo, config *app.Config) {
	studentEp := student.NewStudentEndpoint(config)

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
		{
			Path:     "/grade",
			Method:   http.MethodGet,
			Endpoint: studentEp.CalculateGrade,
		},
		{
			Path:     "/school",
			Method:   http.MethodPost,
			Endpoint: school.AddStudent,
		},
	}

	for _, r := range routes {
		e.Group(r.Group).Add(r.Method, r.Path, r.Endpoint, r.Middleware...)
	}
	// e.GET("/ping", ping.EchoPing)
	// e.POST("/hello", ping.EchoHello)
}
