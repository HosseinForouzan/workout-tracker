package main

import (
	"github.com/HosseinForouzan/workout-tracker.git/auth/authservice"
	"github.com/HosseinForouzan/workout-tracker.git/config"
	"github.com/HosseinForouzan/workout-tracker.git/repository/psql"
	"github.com/HosseinForouzan/workout-tracker.git/user/handler"
	"github.com/HosseinForouzan/workout-tracker.git/user/repository"
	"github.com/HosseinForouzan/workout-tracker.git/user/service"

	ex_handler "github.com/HosseinForouzan/workout-tracker.git/exercise/handler"
	ex_repository "github.com/HosseinForouzan/workout-tracker.git/exercise/repository"
	ex_service "github.com/HosseinForouzan/workout-tracker.git/exercise/service"

	plan_handler "github.com/HosseinForouzan/workout-tracker.git/plan/handler"
	plan_repository "github.com/HosseinForouzan/workout-tracker.git/plan/repository"
	plan_service "github.com/HosseinForouzan/workout-tracker.git/plan/service"

	"github.com/HosseinForouzan/workout-tracker.git/user/validator"
	"github.com/labstack/echo/v5"
)

func main() {

	cfg := config.Load("config.yml")



	psql := psql.New(cfg.Psql)
	userPsql := repository.New(psql)
	exercisePsql := ex_repository.New(psql)
	planPsql := plan_repository.New(psql)
	planExerciseRepo := plan_repository.New(psql)

	userValidator := validator.New(userPsql)

	authSvc := authservice.New(cfg.Auth)
	userSvc := service.New(userPsql, authSvc)
	exerciseSvc := ex_service.New(exercisePsql)
	planSvc := plan_service.New(planPsql, planExerciseRepo)



	e := echo.New()
	
	userHandler := handler.New(cfg, userSvc, userValidator, authSvc)
	userHandler.SetRoutes(e)

	exerciseHandler := ex_handler.New(cfg, exerciseSvc)
	exerciseHandler.SetRoutes(e)

	planHandler := plan_handler.New(cfg, planSvc)
	planHandler.SetRoutes(e)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)

	}	

}