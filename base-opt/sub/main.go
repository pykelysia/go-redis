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

	pubsub := client.Subscribe(ctx, "channel")
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			fmt.Println("ReceiveMessage Error:", err)
			return
		}
		fmt.Println("Received message from channel:", msg.Channel, "Message:", msg.Payload)
	}
}
