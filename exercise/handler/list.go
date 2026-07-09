package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h Handler) exerciseList(c *echo.Context) error {
	resp, err := h.exerciseSvc.ListExercise(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
	
}