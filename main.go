package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jbc2313/goRESTapi/db"
	"github.com/jbc2313/goRESTapi/routes"
)

func helloWorld(c *fiber.Ctx) error {
    return c.SendString("Hello World, from GO!")

}

func setupUserRoutes(app *fiber.App) {
    //User routes
    app.Get("/users", routes.GetUsers)
    app.Post("/users", routes.CreateUser)
    app.Get("/users/:id", routes.GetUser)

}

func main() {
    db.ConnectDb()
    app := fiber.New()

    app.Get("/", helloWorld)

    setupUserRoutes(app)

    log.Fatal(app.Listen(":3050"))
}
