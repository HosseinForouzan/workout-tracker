package handler

import "github.com/labstack/echo/v5"

func (h Handler) SetRoutes(e *echo.Echo) {
	exerciseGroup := e.Group("/exercises")

	exerciseGroup.POST("/add", h.exerciseAdd)
	exerciseGroup.GET("/list", h.exerciseList)

	
}