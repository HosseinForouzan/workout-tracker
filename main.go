package main

import (
	"github.com/HosseinForouzan/workout-tracker.git/config"
	"github.com/HosseinForouzan/workout-tracker.git/repository/psql"
	"github.com/HosseinForouzan/workout-tracker.git/user/handler"
	"github.com/HosseinForouzan/workout-tracker.git/user/repository"
	"github.com/HosseinForouzan/workout-tracker.git/user/service"
	"github.com/HosseinForouzan/workout-tracker.git/user/validator"
	"github.com/labstack/echo/v5"
)

func main() {

	cfg := config.Load("config.yml")

	psql := psql.New(cfg.Psql)
	userPsql := repository.New(psql)

	userSvc := service.New(userPsql)
	userValidator := validator.New(userPsql)

	e := echo.New()
	
	userHandler := handler.New(userSvc, userValidator)
	userHandler.SetRoutes(e)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)

	}	

}