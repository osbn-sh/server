package githubcheckingversionservice

import (
	"context"
	"ostadbun/repository/redis/redisGithubVersionChecking"
)

// ساختار تنظیمات اواث
type GithubCheckingVersionService struct {
	redis redisGithubVersionChecking.GithubVersionChecking
}

func New(redis redisGithubVersionChecking.GithubVersionChecking) *GithubCheckingVersionService {

	return &GithubCheckingVersionService{
		redis: redis,
	}
}

func (o *GithubCheckingVersionService) SetClientVersion(ctx context.Context, version string) error {
	return o.redis.SetClientVersion(ctx, version)
}

func (o *GithubCheckingVersionService) GetClientVersion(ctx context.Context) (string, error) {
	return o.redis.GetClientVersion(ctx)
}

func (o *GithubCheckingVersionService) SetServerVersion(ctx context.Context, version string) error {
	return o.redis.SetServerVersion(ctx, version)
}

func (o *GithubCheckingVersionService) GetServerVersion(ctx context.Context) (string, error) {
	return o.redis.GetServerVersion(ctx)
}
