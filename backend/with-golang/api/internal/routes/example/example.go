package example

import (
	"@dragon-cli-template/apps/api/internal/database"
	"@dragon-cli-template/apps/api/internal/entities"

	"github.com/gofiber/fiber/v2"
)

func ExampleRoutes(app fiber.Router) {
	Example := app.Group("/example")
	Example.Get("/", getExample)
	Example.Post("/", createExample)
}

// Example godoc
// @Summary      Get all Example
// @Tags         Example
// @Accept       json
// @Param        page     query     int  false  "page"   default(1)
// @Param        limit    query     int  false  "limit"  default(10)
// @Param        userId  query     string  true   "userId"
// @Success      200  {object} entities.Example
// @Router       /v1/Example [get]
func getExample(c *fiber.Ctx) error {
	var example []entities.Example

	database.SQL.Find(&example)
	return c.Status(fiber.StatusOK).JSON(example)
}

// Example godoc
// @Summary      Create Example
// @Tags         Example
// @Accept       json
// @Param request body entities.Body true "query params""
// @Success      201  {object} entities.Example
// @Router       /v1/Example [post]
func createExample(c *fiber.Ctx) error {
	example := new(entities.Example)

	database.SQL.Create(&example)
	return c.Status(fiber.StatusCreated).JSON(example)
}
