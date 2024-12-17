package student

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sleepiinuts/kbtg-go-prj/internal/app"
)

type StudentEndpoint struct {
	config *app.Config
}

func NewStudentEndpoint(config *app.Config) *StudentEndpoint {
	return &StudentEndpoint{config: config}
}

func (s *StudentEndpoint) CalculateGrade(c echo.Context) error {
	log.Printf("config score: %d\n", s.config.Score)

	// receive score
	score, err := strconv.Atoi(c.QueryParam("score"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)

	}

	if score < 0 || score > 100 {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("0 <= score <= 100,but got %d", score))
	}

	grade := studentCalculateGrade(score)

	// response score
	return c.JSON(http.StatusOK, grade)

}
