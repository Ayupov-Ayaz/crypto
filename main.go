package main

import (
	"ayupov-ayaz/crypto/crypto/sha256"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

const (
	port = 3000
)

func errorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		return ctx.SendString(err.Error())
	}
}

func main() {
	f := fiber.New(fiber.Config{ErrorHandler: errorHandler()})
	sha256.NewSha256(f)

	if err := f.Listen(":" + strconv.Itoa(port)); err != nil {
		log.Fatal(err)
	}
}
