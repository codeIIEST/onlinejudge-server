package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	router "github.com/raydwaipayan/onlinejudge-server/server/router"
)

func main() {
	app := fiber.New()
	app.Use(recover.New(), logger.New())

	router.SetupRoutes(app)
	app.Listen(":3000")
}
