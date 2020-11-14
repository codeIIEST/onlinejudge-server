package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	config "github.com/raydwaipayan/onlinejudge-server/config"
	router "github.com/raydwaipayan/onlinejudge-server/server/router"
)

func main() {
	app := fiber.New()
	conf, _ := config.Read()

	app.Use(recover.New(), logger.New())

	router.SetupRoutes(app, conf)
	app.Listen(":3000")
}
