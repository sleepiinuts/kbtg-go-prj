package main

import (
	"flag"
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

	option := flag.String("run", "api", "run (api | batch)")
	flag.Parse()

	switch *option {
	case "api":
		echoHTTP(&c)
	default:
		pureHTTP()
	}

}

func pureHTTP() {
	// GET: ping
	http.HandleFunc("/ping", ping.PingHandler)

	http.HandleFunc("/hello", ping.HelloHandler)
	http.ListenAndServe(":8080", nil)
}

func echoHTTP(c *app.Config) {
	e := echo.New()

	// init routes
	handler.InitRoutes(e, c)
	e.Use(handler.TimeConsume)

	e.Start(":8080")
}
