package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	jwtware "github.com/gofiber/jwt/v2"
	handler "github.com/raydwaipayan/onlinejudge-server/server/handler"
)

//SetupRoutes initiates the fiber router
func SetupRoutes(app *fiber.App) {
	godotenv.Load()
	user := app.Group("/user", logger.New(), limiter.New())
	user.Post("/register", handler.Register)
	user.Post("/login", handler.Login)
	user.Use(jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("SECRET_KEY")),
	}))

}
