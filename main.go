package main

import (
    "github.com/gofiber/fiber/v2"
    "log"
    "github.com/jbc2313/goRESTapi/db"

)

func helloWorld(c *fiber.Ctx) error {
    return c.SendString("Hello World, from GO!")

}

func main() {
    db.ConnectDb()
    app := fiber.New()

    app.Get("/", helloWorld)

    log.Fatal(app.Listen(":3050"))
}
