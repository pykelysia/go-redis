package redismq

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func DefultCptClient() *CptClient {
	options := redis.Options(Options{
		Addr: "localhost:6379",
	})
	return &CptClient{
		Client: Client{
			ctx:    context.Background(),
			client: redis.NewClient(&options),
		},
	}
}

func NewPCptClient(options *Options) *CptClient {
	return &CptClient{
		Client: Client{
			ctx:    context.Background(),
			client: redis.NewClient((*redis.Options)(options)),
		},
	}
}
