package handler

import (
	"net/http"

	"github.com/HosseinForouzan/workout-tracker.git/plan/param"
	"github.com/labstack/echo/v5"
)

func (h Handler) planAdd(c *echo.Context) error {
	var req param.PlanAddRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.planSvc.AddPlan(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}