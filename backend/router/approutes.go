package router

import (
	"mailinglist/configs/db"
	"mailinglist/infrastructure/controllers"
	"mailinglist/infrastructure/services"
	"mailinglist/middlewares"

	"github.com/gofiber/fiber/v2"
)

var dbs db.Postgre = db.New()

func start() controllers.Controllers {
	dbs.Validatorpsql()
	var serv services.Services = services.New(dbs)
	var controller controllers.Controllers = controllers.New(serv)
	return controller
}

func Routers(a *fiber.App) {
	controller := start()
	auth := a.Group("/auth")
	auth.Post("/sign-up", controller.SignUp)
	auth.Post("/Sign-in", controller.SignIn)
	secure := auth.Group("/secure").Use(middlewares.Auth())
	secure.Get("/Homepage", controller.Users)
	secure.Delete("/logout", controller.Logout)
	secure.Post("/addcontent", controller.CreateContent)
}
