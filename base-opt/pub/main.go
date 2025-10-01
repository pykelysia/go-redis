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

	err := client.Publish(ctx, "channel", "Hello,GoLang-Redis!").Err()
	if err != nil {
		fmt.Println("Publish err:", err)
		return
	}
	fmt.Println("Publish Success!")
}
