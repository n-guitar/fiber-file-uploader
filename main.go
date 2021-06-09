package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Static file server
	app.Static("/", "./top.html")

	// router
	app.Post("/files", func(c *fiber.Ctx) error {
		// Parse the multipart form:
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		// => *multipart.Form

		files := form.File["document"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

			// Save the files to disk:
			err := c.SaveFile(file, fmt.Sprintf("./files/%s", file.Filename))

			// Check for errors
			if err != nil {
				return err
			}
		}
		return nil
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
