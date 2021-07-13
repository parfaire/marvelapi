package main

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

func InitResources() *redis.Client {
	loadEnv() // has to be the first in the sequence, prior to init other resources
	redisClient := initRedis()

	return redisClient
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}

func initRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic("Cannot connect to redis")
	}

	return redisClient
}
