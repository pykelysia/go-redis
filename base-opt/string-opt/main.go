package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	client.Set(ctx, "key", "value", 60)
	res, err := client.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("key: ", res)
}
