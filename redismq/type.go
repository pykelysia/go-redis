package redismq

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type (
	Client struct {
		// 上下文
		ctx context.Context

		// Redis 客户端
		client *redis.Client
	}

	PSubClient struct {
		Client
	}

	CptClient struct {
		Client
	}

	HandlerFunc func(string, string) bool

	Options redis.Options

	ClientFunc interface {
		Publish(channel string, message interface{}) error
		Subscribe(channel string, handler HandlerFunc) error
	}
)
