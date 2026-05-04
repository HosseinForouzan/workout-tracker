package main

import (
	"fmt"

	"github.com/HosseinForouzan/workout-tracker.git/config"
	"github.com/HosseinForouzan/workout-tracker.git/repository/psql"
)

func main() {

	cfg := config.Load("config.yml")
	// db := psql.New(psql.Config{
	// 	Username: "myuser",
	// 	Password: "secret",
	// 	Port: 5431,
	// 	Host: "localhost",
	// 	DBName: "workout_db",
	// })

	db := psql.New(cfg.Psql)

	fmt.Println(db)
}