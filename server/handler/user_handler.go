package handler

import "github.com/gofiber/fiber/v2"

//Login user login handler
func Login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("passwprd")

	if len(user) == 0 || len(pass) == 0 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.JSON(fiber.Map{"status": "ok"})
}

//Signout user logout handler
func Signout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}
