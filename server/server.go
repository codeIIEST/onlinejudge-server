package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/raydwaipayan/onlinejudge-server/config"
	"github.com/raydwaipayan/onlinejudge-server/server/models"
	router "github.com/raydwaipayan/onlinejudge-server/server/router"
)

func main() {
	app := fiber.New()
	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	app.Use(recover.New(), logger.New())

	router.SetupRoutes(app, conf)
	models.InitDb(conf)
	app.Listen(fmt.Sprintf(":%s", conf.Port))
}
