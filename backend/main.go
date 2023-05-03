package main

import (
	conf "mailinglist/configs"
	"mailinglist/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Server().Concurrency = 100

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		AllowMethods:     "POST, OPTIONS, GET, PUT, DELETE",
	}))

	errs := conf.DBInit()
	if errs != nil {
		panic(errs.Error())
	}

	router.Routers(app)
	app.Listen(":8080")
}
