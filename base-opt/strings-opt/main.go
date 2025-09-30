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
	client.MSet(ctx, map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	})
	res, err := client.MGet(ctx, "key1", "key2").Result()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("key1, key2: ", res)
	fmt.Println("key1: ", res[0])
	fmt.Println("key2: ", res[1])
}
