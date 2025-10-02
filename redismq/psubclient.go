package redismq

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func DefultClient() *PSubClient {
	options := redis.Options(Options{
		Addr: "localhost:6379",
	})
	return &PSubClient{
		Client: Client{
			ctx:    context.Background(),
			client: redis.NewClient(&options),
		},
	}
}

func NewPSubClient(options *Options) *PSubClient {
	return &PSubClient{
		Client: Client{
			ctx:    context.Background(),
			client: redis.NewClient((*redis.Options)(options)),
		},
	}
}
