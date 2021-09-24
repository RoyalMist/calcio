package security

import (
	"time"

	"github.com/gofiber/fiber/v2"
	fiberLimiter "github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimit(limit int, period time.Duration) fiber.Handler {
	return fiberLimiter.New(fiberLimiter.Config{
		Max:        limit,
		Expiration: period,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
	})
}
