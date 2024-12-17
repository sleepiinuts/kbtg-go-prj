package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sleepiinuts/kbtg-go-prj/internal/handler"
	"github.com/sleepiinuts/kbtg-go-prj/internal/ping"
)

func main() {
	echoHTTP()
}

func pureHTTP() {
	// GET: ping
	http.HandleFunc("/ping", ping.PingHandler)

	http.HandleFunc("/hello", ping.HelloHandler)
	http.ListenAndServe(":8080", nil)
}

func echoHTTP() {
	e := echo.New()

	// init routes
	handler.InitRoutes(e)
	e.Use(handler.TimeConsume)

	e.Start(":8080")
}
