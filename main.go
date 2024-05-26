package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Response struct {
	Page string
	Data interface{}
}

func (r Response) PageIs(page string) bool {
	return r.Page == page
}

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", Response{
			Page: "Home",
			Data: fiber.Map{
				"name": "jon snow!",
			},
		}, "layouts/main")
	})

	app.Get("/genres", func(c *fiber.Ctx) error {
		return c.Render("genres", Response{
			Page: "Genres",
			Data: fiber.Map{
				"genres": map[int]interface{}{
					1: "Action",
					2: "Adventure",
					3: "Comedy",
				},
			},
		}, "layouts/main")
	})

	app.Get("/movies", func(c *fiber.Ctx) error {
		return c.Render("movies", Response{
			Page: "Movies",
			Data: fiber.Map{
				"movies": map[int]interface{}{
					1: "The Shawshank Redemption",
					2: "The Godfather",
					3: "The Dark Knight",
				},
			},
		}, "layouts/main")
	})

	app.Post("/count", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("count", fiber.Map{
			"count": 1,
		})
	})

	log.Fatal(app.Listen(":1323"))
}
