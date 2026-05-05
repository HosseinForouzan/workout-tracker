package claim

import (
	"github.com/HosseinForouzan/workout-tracker.git/auth/authservice"
	"github.com/HosseinForouzan/workout-tracker.git/config"
	"github.com/labstack/echo/v5"
)

func GetClaimsFromEchoContext(c *echo.Context) *authservice.Claims {
	return c.Get(config.AuthMiddlewareContextKey).(*authservice.Claims)
}