package main

import (
    "context"
    "github.com/gofiber/fiber/v2"
    "github.com/go-redis/redis/v8"
    "log"
    "time"
)

var ctx = context.Background()

func main() {
    app := fiber.New()

    // Redis client setup
    rdb := redis.NewClient(&redis.Options{
        Addr: "redis:6379",
        Password: "", // no password set
        DB: 0,  // use default DB
    })

    // Middleware to cache responses
    app.Use(func(c *fiber.Ctx) error {
        key := "cache:" + c.Path()

        // Check if the cache exists
        cached, err := rdb.Get(ctx, key).Result()
        if err == redis.Nil {
            // Cache miss: Proceed with the request
            if err := c.Next(); err != nil {
                return err
            }

            // Store the response in Redis with a TTL of 1 minute
            rdb.Set(ctx, key, c.Response().Body(), time.Minute)

            return nil
        } else if err != nil {
            return c.Status(500).SendString("Redis error: " + err.Error())
        }

        // Cache hit: Return the cached response
        return c.SendString(cached)
    })

    // Example route
    app.Get("/books", func(c *fiber.Ctx) error {
        // Assume this is a costly operation that we want to cache
        books := "List of books"
        return c.SendString(books)
    })

    log.Fatal(app.Listen(":8080"))
}
