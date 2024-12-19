package school

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sleepiinuts/kbtg-go-prj/internal/model"
)

type Endpoint struct{}

type req struct {
	Room  int    `json:"room"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func AddStudent(c echo.Context) error {
	var req req
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	// business logic
	// student
	s := model.Student{
		Name:  req.Name,
		Score: req.Score,
	}

	if err := addStudent(req.Room, s); err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusCreated, nil)
}
