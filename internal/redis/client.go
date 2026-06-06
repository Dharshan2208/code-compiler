package redis

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	goredis "github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func New() *goredis.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}

	client := goredis.NewClient(&goredis.Options{
		Addr: addr,
	})

	if err := client.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Printf("Connected to Redis at %s", addr)

	return client
}
