package main

import (
	db "@dragon-cli-template/apps/api/internal/database"
	"@dragon-cli-template/apps/api/internal/pkg/env"
	"@dragon-cli-template/apps/api/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	db.Open()
	routes.InitRoutes(app)

	app.Listen(":" + env.Envs.Port)
}
