package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raydwaipayan/onlinejudge-server/server/models"
	"github.com/raydwaipayan/onlinejudge-server/server/types"
)

// CreateContest Creates new contest
func CreateContest(c *fiber.Ctx) error {
	u := new(types.Contest)
	if err := c.BodyParser(u); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err := u.ContestCreate(models.DBConfigURL)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(u)
}
