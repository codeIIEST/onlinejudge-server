package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raydwaipayan/onlinejudge-server/server/middlewares"
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
	return c.SendStatus(fiber.StatusCreated)
}

// RegisterContest Registers user for contest
func RegisterContest(c *fiber.Ctx) error {
	contest := new(types.Contest)
	handle := new(types.UserHandle)
	cid := c.Params("cid")
	contest.ID = cid
	found, _ := contest.CheckContestExists(models.DBConfigURL)

	if !found {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if err := c.BodyParser(handle); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if len(handle.Handle) == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	contest.Users = middlewares.UniqueHandles(append(contest.Users, handle.Handle))

	err := contest.UpdateContest(models.DBConfigURL)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(contest)
}

// GetAllContests Gets all contest
func GetAllContests(c *fiber.Ctx) error {
	contest := new(types.Contest)

	allContests, err := contest.GetAllContests(models.DBConfigURL)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(allContests)
}
