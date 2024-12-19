package school

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sleepiinuts/kbtg-go-prj/internal/model"
)

type Endpoint struct {
	Serv Schooler
}

type Schooler interface {
	AddStudentToDB(room int, stu model.Student) error
	GetStudentByRoom(room int) ([]model.Student, error)
}

type req struct {
	Room  int    `json:"room"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (e *Endpoint) AddStudent(c echo.Context) error {
	var req req
	if err := c.Bind(&req); err != nil {
		// return echo.ErrBadRequest
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// business logic
	// student
	s := model.Student{
		Name:  req.Name,
		Score: req.Score,
	}

	if err := e.Serv.AddStudentToDB(req.Room, s); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, nil)
}

func (e *Endpoint) ListStudent(c echo.Context) error {
	room, err := strconv.Atoi(c.QueryParam("room"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// business logic
	students, err := e.Serv.GetStudentByRoom(room)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, students)
}

func (e *Endpoint) CountStudentInRoom(c echo.Context) error {
	room, err := strconv.Atoi(c.QueryParam("room"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// business logic
	students, err := e.Serv.GetStudentByRoom(room)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, len(students))
}
