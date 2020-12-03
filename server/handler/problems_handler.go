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
	contest.ID = cid
	found, _ := contest.CheckContestExists(models.DBConfigURL)

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
	contest.Problems = append(contest.Problems, problem.Index)
	updateErr := contest.UpdateContest(models.DBConfigURL)
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
	notFound, err := problem.GetProblem(models.DBConfigURL)

	if notFound {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(problem)
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
