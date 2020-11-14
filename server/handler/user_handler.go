package handler

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/raydwaipayan/onlinejudge-server/config"
	"github.com/raydwaipayan/onlinejudge-server/server/types"
)

//Register types.User registration handler
func Register(c *fiber.Ctx) error {
	u := new(types.User)

	if err := c.BodyParser(u); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

//Login types.User login handler
func Login(conf *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u := new(types.User)

		if err := c.BodyParser(u); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = u.Email
		claims["admin"] = false
		claims["exp"] = time.Now().Add(time.Hour * 96).Unix()

		t, err := token.SignedString([]byte(conf.SecretKey))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(fiber.Map{"token": t})
	}
}
