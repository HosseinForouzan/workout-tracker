package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	e.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, "health ok")
	})
	userGroup := e.Group("/users")

	userGroup.POST("/register", h.userRegister)
	userGroup.POST("/login", h.userLogin)
}