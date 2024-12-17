package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func TimeConsume(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		next(c)
		stop := time.Now()

		log.Println("time consume by middleware: ", stop.Sub(start))
		return nil
	}
}

func MovePermanant(_ echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.JSON(http.StatusMovedPermanently, nil)
		return nil
	}
}
