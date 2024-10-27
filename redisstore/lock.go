package redisstore

import (
	"context"
	"time"

	"github.com/li-zeyuan/common-go/mylogger"
	"go.uber.org/zap"
)

func (r *RedisDb) Lock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	isSet, err := r.Cli.SetNX(ctx, key, 1, expiration).Result()
	if err != nil {
		mylogger.Error(ctx, "redis set lock fail", zap.Error(err), zap.String("key", key))
		return false, err
	}

	return isSet, nil
}

func (r *RedisDb) UnLock(ctx context.Context, key string) error {
	err := r.Cli.Del(ctx, key).Err()
	if err != nil {
		mylogger.Error(ctx, "redis unlock fail", zap.Error(err), zap.String("key", key))
		return err
	}

	return nil
}
