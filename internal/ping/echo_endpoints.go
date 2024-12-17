package ping

import (
	"net/http"

	"github.com/labstack/echo"
)

func EchoPing(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}

func EchoHello(ctx echo.Context) error {
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
}
