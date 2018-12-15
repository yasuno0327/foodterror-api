package service

import (
	"github.com/go-redis/redis"
)

var client *redis.Client

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func SetSession(userId uint) (string, error) {
	key, err := GenerateToken()
	if err != nil {
		return "", err
	}
	if err = client.Set(key, userId, 0).Err(); err != nil {
		return "", err
	}
	return key, nil
}
