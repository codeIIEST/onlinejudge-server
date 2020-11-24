package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raydwaipayan/onlinejudge-server/server/models"
	"github.com/raydwaipayan/onlinejudge-server/server/types"
)

// CreateProblem creates a new problem based on contest id
func CreateProblem(c *fiber.Ctx) error {

	cid := c.Params("cid")
	contest := new(types.Contest)
	found, _ := contest.CheckContestExists(models.DBConfigURL, cid)

	if !found {
		return c.SendStatus(fiber.StatusNotFound)
	}

	problem := new(types.Problem)

	if err := c.BodyParser(problem); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	problem.ContestID = cid
	exists, err := problem.CreateProblem(models.DBConfigURL)
	if exists {
		return c.SendStatus(fiber.StatusConflict)
	}
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusCreated)
}
