package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong!")
	})

	posts := [4]Post{
		{
			"Daerah Paling Ramah di Amerika?",
			"Michi-gan",
		},
		{
			"Makanan Favorit Batman?",
			"Batagor",
		},
		{
			"Apa beda nya Soto sama Cotto?",
			"Soto pakai daging sapi, kalau Cotto pakai daging Capii",
		},
		{
			"Soto khas Jepang?",
			"Chotto Matte",
		},
	}

	app.Get("/recents", func(c *fiber.Ctx) error {
		latestPosts := [4]Post{posts[3], posts[2], posts[1], posts[0]}

		return c.Render("recents", fiber.Map{
			"Jokes": latestPosts,
		})
	})

	app.Get("/post/:id", func(c *fiber.Ctx) error {
		postId, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			c.SendString("Not Found")
			return c.SendStatus(404)
		}

		if postId < 0 || postId >= len(posts) {
			c.SendString("Not Found")
			return c.SendStatus(404)
		}
		post := posts[postId]
		return c.Render("detail", fiber.Map{
			"Joke": post,
		})
	})

	app.Static("/", "./public")
	// start server
	log.Fatal(app.Listen(":8000"))
}
