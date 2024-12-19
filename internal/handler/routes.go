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

	var serv *school.Service
	switch config.Database {
	case "mongo":
		serv = school.NewService(&school.RedisRepos{})
	case "redis":
		serv = school.NewService(&school.MongoRepos{})
	default:
		serv = school.NewService(&school.MongoRepos{})
	}

	// school
	// rp := school.RedisRepos{}
	// schoolServ := school.NewService(&rp)
	schoolEp := school.Endpoint{
		Serv: serv,
	}

	// mongoRp := school.MongoRepos{}
	// schoolMongoServ := school.NewService(&mongoRp)
	schoolMongoEp := school.Endpoint{
		Serv: serv,
	}

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
			Endpoint: schoolEp.AddStudent,
		},
		{
			Path:     "/schoolMongo",
			Method:   http.MethodPost,
			Endpoint: schoolMongoEp.AddStudent,
		},
		{
			Group:    "mongo",
			Path:     "/school-listStudent",
			Method:   http.MethodGet,
			Endpoint: schoolMongoEp.ListStudent,
		},
		{
			Group:    "mongo",
			Path:     "/school-countStudent",
			Method:   http.MethodGet,
			Endpoint: schoolMongoEp.CountStudentInRoom,
		},
	}
	for _, r := range routes {
		e.Group(r.Group).Add(r.Method, r.Path, r.Endpoint, r.Middleware...)
	}
	// e.GET("/ping", ping.EchoPing)
	// e.POST("/hello", ping.EchoHello)
}
