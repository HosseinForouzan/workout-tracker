package handler

import (
	"net/http"

	"github.com/HosseinForouzan/workout-tracker.git/exercise/param"
	"github.com/labstack/echo/v5"
)

func (h Handler) exerciseAdd(c *echo.Context) error {
	var req param.ExerciseRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.exerciseSvc.AddExercise(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "Exercise Created")
}