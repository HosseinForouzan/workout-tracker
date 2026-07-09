package handler

import "github.com/labstack/echo/v5"

func (h Handler) SetRoutes(e *echo.Echo) {
	planGroup := e.Group("/plans")

	planGroup.POST("", h.planAdd)
	planGroup.GET("", h.getAll)
	planGroup.GET("/:id", h.getUserPlans)

	planGroup.POST("/exercises", h.planExerciseAdd)


	
}