package repo_redis

import "github.com/redis/go-redis/v9"

type sceneRedisRepo struct {
	redisClient *redis.Client
}

type SceneRedisRepo interface {
}

func NewSceneRedisRepo(redisClient *redis.Client) SceneRedisRepo {
	return &sceneRedisRepo{
		redisClient: redisClient,
	}
}
