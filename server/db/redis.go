package db

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v7"
)

// Singleton to keep assets loaded in memory
type redisManager struct {
	client *redis.Client
}

var singleton *redisManager
var once sync.Once

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetDB() *redisManager {
	once.Do(func() {
		redis_port, err := strconv.Atoi(getEnv("REDIS_PORT", "6379"))
		if err != nil {
			panic("Invalid REDIS_PORT specified")
		}
		redis_db, err := strconv.Atoi(getEnv("REDIS_DB", "0"))
		if err != nil {
			panic("Invalid REDIS_DB specified")
		}
		client := redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", getEnv("REDIS_HOST", "localhost"), redis_port),
			DB:   redis_db,
		})
		// Create object
		singleton = &redisManager{
			client: client,
		}
	})
	return singleton
}

// Get - Redis GET
func (r *redisManager) Get(key string) string {
	return r.Get(key)
}

// Get - Redis SET
func (r *redisManager) Set(key string, value string) string {
	return r.Set(key, value)
}
