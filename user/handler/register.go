package handler

import (
	"net/http"

	"github.com/HosseinForouzan/workout-tracker.git/user/param"
	"github.com/labstack/echo/v5"
)

func (h Handler) userRegister(c *echo.Context) error {
	var req param.RegisterReqeust
	
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if fieldErros, err := h.userValidator.ValidateRegisterRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"errors": fieldErros})
	}

	resp, err := h.userSvc.Register(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}