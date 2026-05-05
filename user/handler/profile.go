package handler

import (
	"net/http"

	"github.com/HosseinForouzan/workout-tracker.git/pkg/claim"
	"github.com/HosseinForouzan/workout-tracker.git/user/param"
	"github.com/labstack/echo/v5"
)

func (h Handler) userProfile(c *echo.Context) error {
	claims := claim.GetClaimsFromEchoContext(c)

	resp, err := h.userSvc.Profile(c.Request().Context(), param.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)

}