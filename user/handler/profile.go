package handler

import (
	"net/http"
	"strconv"

	"github.com/HosseinForouzan/workout-tracker.git/user/param"
	"github.com/labstack/echo/v5"
)

func (h Handler) userProfile(c *echo.Context) error {
	id , _:= strconv.Atoi(c.Param("id"))
	resp, err := h.userSvc.Profile(c.Request().Context(), param.ProfileRequest{UserID: uint(id)})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)

}