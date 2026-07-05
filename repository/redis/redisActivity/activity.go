package redisActivity

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisActivity struct {
	redis *redis.Client
}

func New(client *redis.Client) *RedisActivity {
	return &RedisActivity{redis: client}
}

func (o *RedisActivity) Set(ctx context.Context, userid, level int) error {
	fmt.Println(userid, level)

	key := fmt.Sprintf("user-level:%d", userid)

	return o.redis.Set(ctx, key, level, time.Hour*6).Err()
}

func (o *RedisActivity) Check(ctx context.Context, userid int) (int, error) {
	key := fmt.Sprintf("user-level:%d", userid)

	res, err := o.redis.Get(ctx, key).Result()

	if errors.Is(err, redis.Nil) {
		return -1, nil
	}

	if err != nil {
		return -1, err
	}

	level, err := strconv.Atoi(res)
	if err != nil {

		return -1, err
	}

	return level, nil
}
