package config

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var SessionStore *session.Store

func InitSessionStore() error {
	storage := redis.New(redis.Config{
		URL:   os.Getenv("REDIS_URL"),
		Reset: false,
	})

	SessionStore = session.New(session.Config{
		Storage:        storage,
		Expiration:     24 * time.Hour,
		CookiePath:     "/",
		CookieHTTPOnly: true,
	})

	return nil
}
