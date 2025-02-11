package main

import (
	"github.com/alexndr54/library-buyer-digiflazz/webook"
	"github.com/gofiber/fiber/v2"
)

func main() {
	f := fiber.New()

	f.Post("/", func(ctx *fiber.Ctx) error {
		err, p := webook.GofiberWebhookHandler(ctx)
		if err != nil {
			return ctx.SendString(err.Error())
		}

		return ctx.JSON(p)
	})

	err := f.Listen(":1000")
	if err != nil {
		panic(err.Error())
	}
}
