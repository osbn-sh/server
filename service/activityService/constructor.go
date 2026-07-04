package activityService

import (
	"ostadbun/repository/postgres/activityRepository"
	"ostadbun/repository/redis/redisActivity"
)

type Activity struct {
	repo  *activityRepository.DB
	redis redisActivity.RedisActivity
}

func New(repo *activityRepository.DB, redis redisActivity.RedisActivity) Activity {
	return Activity{
		repo:  repo,
		redis: redis,
	}
}
