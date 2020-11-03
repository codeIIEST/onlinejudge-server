package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"

	jwtware "github.com/gofiber/jwt/v2"
	handler "github.com/raydwaipayan/onlinejudge-server/server/handler"
)

//SetupRoutes initiates the fiber router
func SetupRoutes(app *fiber.App) {
	user := app.Group("/user", logger.New(), limiter.New())

	user.Post("/login", handler.Login)
	user.Use(jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    "test",
	}))

	user.Get("/signout", handler.Signout)
}
