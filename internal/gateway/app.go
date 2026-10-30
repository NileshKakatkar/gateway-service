package gateway

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func NewApp(cfg Config) *fiber.App {
	client := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})
	app := fiber.New(fiber.Config{AppName: "ShopCore API Gateway", BodyLimit: 2 << 20})
	app.Use(recover.New())
	app.Get("/health", func(c *fiber.Ctx) error {
		if err := client.Ping(context.Background()).Err(); err != nil {
			return fiber.NewError(503, "rate limiter unavailable")
		}
		return c.JSON(fiber.Map{"status": "ok", "service": "gateway"})
	})
	app.Use("/api", func(c *fiber.Ctx) error {
		key := "gateway:" + c.IP()
		count, err := client.Incr(context.Background(), key).Result()
		if err != nil {
			return fiber.NewError(503, "rate limiter unavailable")
		}
		if count == 1 {
			if err := client.Expire(context.Background(), key, time.Minute).Err(); err != nil {
				return fiber.NewError(503, "rate limiter unavailable")
			}
		}
		if err == nil && count > 120 {
			return fiber.NewError(429, "rate limit exceeded")
		}
		c.Set("X-Request-ID", strconv.FormatInt(time.Now().UnixNano(), 10))
		return c.Next()
	})
	app.All("/api/v1/auth/*", func(c *fiber.Ctx) error { return proxy.Do(c, cfg.UserURL+c.OriginalURL()) })
	app.All("/api/v1/users/*", func(c *fiber.Ctx) error { return proxy.Do(c, cfg.UserURL+c.OriginalURL()) })
	app.All("/api/v1/categories*", func(c *fiber.Ctx) error { return proxy.Do(c, cfg.ProductURL+c.OriginalURL()) })
	app.All("/api/v1/products*", func(c *fiber.Ctx) error { return proxy.Do(c, cfg.ProductURL+c.OriginalURL()) })
	app.All("/api/v1/orders*", func(c *fiber.Ctx) error { return proxy.Do(c, cfg.OrderURL+c.OriginalURL()) })
	app.Use(func(c *fiber.Ctx) error { return fiber.NewError(404, "route not found") })
	return app
}
