package main

import (
	"net/http"

	"github.com/labstack/echo"
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
	e.GET("/ping", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "pong")
	})

	e.Start(":8080")
}
