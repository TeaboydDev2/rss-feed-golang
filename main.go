package main

import (
	"rss-sandbox/rss_model"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	app.Static("/rss/t", "./rss_model/feed.xml")
	app.Get("/rss-news", rss_model.ServeRss)
	app.Get("/user-rss", rss_model.RssGenerator)
	app.Listen(":2500")
}
