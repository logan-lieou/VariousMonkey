package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func params(c *fiber.Ctx) error {
	return c.SendString("value: " + c.Params("value"))
}

func op_params(c *fiber.Ctx) error {
	if c.Params("name") != "" {
		return c.SendString("Hello " + c.Params("name"))
	}
	return c.SendString("What's the name again?")
}

func main() {
	// init our application
	app := fiber.New()
	app.Use(cors.New())

	// simple route
	app.Get("/", hello)

	// route with parameters
	app.Get("/:value", params)

	// route with optinal parameters
	app.Get("/:name?", op_params)

	// static site server
	app.Static("/static", "./static")

	// using middleware with minimal config
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "12345",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "john" && pass == "doe" {
				return true
			}
			if user == "admin" && pass == "12345" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.SendFile("./static/401.html")
		},
		ContextUsername: "_user",
		ContextPassword: "_password",
	}))

	app.Listen(":3000")
}
