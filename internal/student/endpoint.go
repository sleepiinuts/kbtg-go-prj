package student

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CalculateGrade(c echo.Context) error {
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
