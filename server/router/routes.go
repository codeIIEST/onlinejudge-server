package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	jwtware "github.com/gofiber/jwt/v2"
	config "github.com/raydwaipayan/onlinejudge-server/config"
	handler "github.com/raydwaipayan/onlinejudge-server/server/handler"
)

// SetupRoutes initiates the fiber router
func SetupRoutes(app *fiber.App, conf *config.Config) {
	api := app.Group("/api/v1", limiter.New())

	user := api.Group("/user")
	contest := api.Group("/contest")
	problems := api.Group("/problem")

	user.Post("/register", handler.Register)
	user.Post("/login", handler.Login(conf))

	contest.Get("/all", handler.GetAllContests)

	problems.Get("/:cid/all", handler.GetAllProblems)
	problems.Get("/:index", handler.GetProblem)

	api.Use(jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(conf.SecretKey),
	}))

	contest.Post("/create", handler.CreateContest)
	contest.Post("/:cid/register", handler.RegisterContest)

	problems.Post("/:cid/submit", handler.CreateProblem)

}
