package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	log.Printf("--> %s %s", c.Method(), c.OriginalURL())
	log.Printf("Headers: %v", c.GetReqHeaders())
	log.Printf("Query: %v", c.Queries())
	log.Printf("Body: %v", c.Body())

	err := c.Next()

	duration := time.Since(start)
	log.Printf("Request to %s took %v", c.OriginalURL(), duration)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	return err
}
