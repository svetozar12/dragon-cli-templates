package routes

import (
	"@dragon-cli-template/config"
	content "@dragon-cli-template/routes/content"
	contentmodel "@dragon-cli-template/routes/contentModel"
	"@dragon-cli-template/routes/field"
	fieldtype "@dragon-cli-template/routes/fieldType"

	_ "@dragon-cli-template/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:4000
// @BasePath /
func InitRoutes(app *fiber.App) {
	config.Init()
	v1 := app.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from the api")
	})
	v1.Get("/swagger/*", swagger.HandlerDefault) // default
	fieldtype.FieldTypeRoutes(v1)
	contentmodel.ContentModelRoutes(v1)
	content.ContentRoutes(v1)
	field.FieldRoutes(v1)
}
