package redisGithubVersionChecking

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type GithubVersionChecking struct {
	redis *redis.Client
}

func New(client *redis.Client) *GithubVersionChecking {
	return &GithubVersionChecking{redis: client}
}

const (
	clientVersionKey = "github:client:version"
	serverVersionKey = "github:server:version"
)

func (o *GithubVersionChecking) SetClientVersion(ctx context.Context, version string) error {
	return o.redis.Set(ctx, clientVersionKey, version, 0).Err()
}

func (o *GithubVersionChecking) GetClientVersion(ctx context.Context) (string, error) {
	return o.redis.Get(ctx, clientVersionKey).Result()
}

func (o *GithubVersionChecking) SetServerVersion(ctx context.Context, version string) error {
	return o.redis.Set(ctx, serverVersionKey, version, 0).Err()
}

func (o *GithubVersionChecking) GetServerVersion(ctx context.Context) (string, error) {
	return o.redis.Get(ctx, serverVersionKey).Result()
}
