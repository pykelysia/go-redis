package redismq

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type (
	Client struct {
		ctx    context.Context // 上下文
		client *redis.Client   // Redis 客户端
	}

	HandlerFunc func(string, ...string) bool
)

func DefultClient() *Client {
	return &Client{
		ctx: context.Background(),
		client: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		}),
	}
}
