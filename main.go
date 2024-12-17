package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sleepiinuts/kbtg-go-prj/internal/app"
	"github.com/sleepiinuts/kbtg-go-prj/internal/handler"
	"github.com/sleepiinuts/kbtg-go-prj/internal/ping"
)

func main() {

	c := app.Config{}
	if err := c.Init("dev"); err != nil {
		log.Printf("err: %v\n", err)
	}

	log.Printf("config: %+v\n", c)
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
