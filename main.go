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
	e.POST("/hello", func(ctx echo.Context) error {
		var req struct {
			Name    string `json:"name"`
			Surname string `json:"surname"`
			Address struct {
				No   int    `json:"no"`
				Road string `json:"road"`
			} `json:"address"`
		}

		if err := ctx.Bind(&req); err != nil {
			return ctx.String(http.StatusBadRequest, "Bad Request")
		}

		return ctx.JSON(http.StatusOK, map[string]any{"message": "success", "response": req})
	})

	e.Start(":8080")
}
