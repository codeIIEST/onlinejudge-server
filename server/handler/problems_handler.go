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
	contestData, found, _ := contest.CheckContestExists(models.DBConfigURL, cid)

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
	// Appends problem to contest problems
	contestData.Problems = append(contestData.Problems, problem.Index)
	updateErr := contestData.UpdateContest(models.DBConfigURL, cid)
	if updateErr != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusCreated)
}

// GetProblem gets problem based on index
func GetProblem(c *fiber.Ctx) error {

	index := c.Params("index")
	problem := new(types.Problem)
	problem.Index = index
	data, notFound, err := problem.GetProblem(models.DBConfigURL)

	if err != nil {
		if notFound {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(data)
}

// GetAllProblems gets all problems for particular cid
func GetAllProblems(c *fiber.Ctx) error {

	problem := new(types.Problem)
	cid := c.Params("cid")
	data, err := problem.GetAllProblems(models.DBConfigURL, cid)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if len(*data) == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(data)
}
