package ping

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sleepiinuts/kbtg-go-prj/internal/model"
)

func EchoPing(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}

func EchoHello(ctx echo.Context) error {
	var req model.Request

	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad Request")
	}

	return ctx.JSON(http.StatusOK, map[string]any{"message": "success", "response": req})
}
