package handler

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/raydwaipayan/onlinejudge-server/config"
	"github.com/raydwaipayan/onlinejudge-server/server/models"
	"github.com/raydwaipayan/onlinejudge-server/server/types"
	"golang.org/x/crypto/bcrypt"
)

// Register types.User registration handler
func Register(c *fiber.Ctx) error {
	u := new(types.User)

	if err := c.BodyParser(u); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	pass := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	u.Password = string(hashedPassword)
	err = u.Create(models.DBConfigURL)
	if  err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}

// Login types.User login handler
// Use closures to handle parameters
func Login(conf *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u := new(types.User)

		if err := c.BodyParser(u); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		doesUserExist, _ := u.CheckUserExists(models.DBConfigURL)

		if doesUserExist == false {
			log.Println("User has not registered")
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = u.Email
		claims["exp"] = time.Now().Add(time.Hour * 96).Unix()

		t, err := token.SignedString([]byte(conf.SecretKey))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(fiber.Map{"token": t})
	}
}
