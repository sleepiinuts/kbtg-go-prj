package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func TimeConsume(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		err := next(c)

		// handle interserver error
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
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
