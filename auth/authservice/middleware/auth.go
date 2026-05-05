package middleware

import (
	"github.com/HosseinForouzan/workout-tracker.git/auth/authservice"
	cfg "github.com/HosseinForouzan/workout-tracker.git/config"
	mw "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

func Auth(service authservice.Service, config authservice.Config) echo.MiddlewareFunc {
	return mw.WithConfig(mw.Config{
		ContextKey:cfg.AuthMiddlewareContextKey ,
		SigningKey: []byte(config.SignKey),
		SigningMethod: "HS256",
		ParseTokenFunc: func(c *echo.Context, auth string) (interface{}, error) {
			claims, err := service.ParseToken(auth)
			if err != nil {
				return nil, err
			}

			return claims, nil
		},
	})
}