package redisstore

import (
	"context"
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	ctx := context.Background()
	rdb, err := New(context.Background(), &Config{
		DSN: "redis://:@localhost:6379/1",
	})
	if err != nil {
		t.Fatal(err)
	}

	err = rdb.Cli.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Cli.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
