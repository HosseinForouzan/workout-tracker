package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func (h Handler) getAll(c *echo.Context) error {
	resp, err := h.planSvc.GetAllPlans(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Handler) getUserPlans(c *echo.Context) error {
	userID := c.Param("id")
	userIdUint, _ := strconv.Atoi(userID)
	resp, err := h.planSvc.GetPlansOfUser(c.Request().Context(), uint(userIdUint))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}