package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/digiconvent/qap/demo/server/realm"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	s := &server{realm: realm.New()}
	app := fiber.New()

	app.Post("/demo/reset", s.reset)
	app.Get("/demo/record", s.record)

	app.Post("/api/shire/bag-end/bequeath", s.requires("shire.bag-end.bequeath"))
	app.Post("/api/council/ring/decide", s.requires("council.ring.decide"))
	app.Post("/api/council/fellowship/form", s.requires("council.fellowship.form"))
	app.Post("/api/fellowship/ring/bearer", s.requires("fellowship.ring.bearer"))
	app.Post("/api/isengard/peace/offer", s.requires("isengard.peace.offer"))
	app.Post("/api/fellowship/journey/plan", s.requires("fellowship.journey.plan"))
	app.Post("/api/rohan/court/grant-authority", s.requires("rohan.court.grant-authority"))
	app.Post("/api/rohan/court/banish", s.requires("rohan.court.banish"))
	app.Post("/api/fangorn/isengard/march", s.requires("fangorn.isengard.march"))
	app.Get("/api/shire/red-book", s.requires("shire.red-book.read"))

	app.Get("/*", static.New(filepath.Join("..", "web")))

	port := os.Getenv("PORT")
	if port == "" {
		port = "4280"
	}
	log.Fatal(app.Listen(":" + port))
}
