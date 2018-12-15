package service

import (
	"strconv"

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

func GetUserId(key string) (int, error) {
	userId, err := client.Get(key).Result()
	if err != nil {
		return 0, err
	}
	id, err := strconv.Atoi(userId)
	if err != nil {
		return 0, err
	}
	return id, err
}
