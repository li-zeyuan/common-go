package redisstore

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRedisDb_Lock(t *testing.T) {
	ctx := context.Background()
	rdb, err := New(ctx, &Config{
		DSN:                       "redis://:@localhost:6379/1",
		DefaultUnlockTimeDuration: 0,
	})
	if err != nil {
		t.Fatal(err)
	}

	isSet, err := rdb.Lock(ctx, "111", time.Second*50)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, isSet)

	isSet, err = rdb.Lock(ctx, "111", time.Second*100)
	if err != nil {
		t.Fatal(err)
	}
	assert.False(t, isSet)
}

func TestUnLock(t *testing.T) {
	ctx := context.Background()
	rdb, err := New(ctx, &Config{
		DSN:                       "redis://:@localhost:6379/1",
		DefaultUnlockTimeDuration: 0,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = rdb.UnLock(ctx, "111")
	assert.Nil(t, err)
}
