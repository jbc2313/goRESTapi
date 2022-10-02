package main

import (
    "github.com/gofiber/fiber/v2"
    "log"

)

func helloWorld(c *fiber.Ctx) error {
    return c.SendString("Hello World, from GO!")

}

func main() {
    app := fiber.New()

    app.Get("/", helloWorld)

    log.Fatal(app.Listen(":3050"))
}
