package config

import "github.com/HosseinForouzan/workout-tracker.git/repository/psql"

type Config struct {
	Psql psql.Config `koanf:"psql"`
}