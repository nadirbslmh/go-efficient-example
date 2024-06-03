package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// use rate limiter
	rateLimiterConfig := middleware.RateLimiterMemoryStoreConfig{
		Rate:      10,
		Burst:     30,
		ExpiresIn: 3 * time.Minute,
	}

	memoryStore := middleware.NewRateLimiterMemoryStoreWithConfig(rateLimiterConfig)
	rateLimiterMiddleware := middleware.RateLimiter(memoryStore)

	e.Use(rateLimiterMiddleware)
	e.Use(middleware.Logger())

	// routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
