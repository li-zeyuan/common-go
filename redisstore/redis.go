package redisstore

import (
	"context"
	"errors"

	"github.com/li-zeyuan/common-go/mylogger"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisDb struct {
	conf *Config
	Cli  *redis.Client
}

var _redis *RedisDb

func New(ctx context.Context, cfg *Config) (*RedisDb, error) {
	if _redis != nil {
		return _redis, nil
	}

	opts, err := redis.ParseURL(cfg.DSN)
	if err != nil {
		mylogger.Error(ctx, "redis parse url fail", zap.Error(err))
		return nil, err
	}

	cli := redis.NewClient(opts)
	if err = cli.Ping(ctx).Err(); err != nil {
		mylogger.Error(ctx, "redis ping fail", zap.Error(err))
		return nil, err
	}

	_redis = &RedisDb{
		conf: cfg,
		Cli:  cli,
	}

	return _redis, nil
}

func GetClient(ctx context.Context) (*RedisDb, error) {
	if _redis != nil {
		return _redis, nil
	}

	mylogger.Error(ctx, "redis client uninitialized")
	return nil, errors.New("redis client uninitialized")
}

func (r *RedisDb) Close() {
	r.Cli.Close()
}
