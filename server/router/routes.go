package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"

	jwtware "github.com/gofiber/jwt/v2"
	config "github.com/raydwaipayan/onlinejudge-server/config"
	handler "github.com/raydwaipayan/onlinejudge-server/server/handler"
)

//SetupRoutes initiates the fiber router
func SetupRoutes(app *fiber.App, conf *config.Config) {
	user := app.Group("/user", logger.New(), limiter.New())
	user.Post("/register", handler.Register)
	user.Post("/login", handler.Login(conf))
	user.Use(jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(conf.SecretKey),
	}))

}
