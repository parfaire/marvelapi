package character_controller

import (
	"github.com/go-redis/redis"
	"github.com/parfaire/marvelapi/channels/marvel"
	"github.com/parfaire/marvelapi/controllers"
)

type CharacterController struct {
	controllers.Controller
	RedisClient   *redis.Client
	MarvelChannel marvel.Interface
}

func New(redisClient *redis.Client, marvelChannel marvel.Interface) CharacterController {
	return CharacterController{
		RedisClient:   redisClient,
		MarvelChannel: marvelChannel,
	}
}
