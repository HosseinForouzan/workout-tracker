package handler

import (
	"net/http"

	"github.com/HosseinForouzan/workout-tracker.git/user/param"
	"github.com/labstack/echo/v5"
)

func (h Handler) userLogin(c *echo.Context) error {
	var req param.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}


	if fieldErrors, err := h.userValidator.ValidateLoginRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"errors": fieldErrors})
	}

	resp, err := h.userSvc.Login(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)


}