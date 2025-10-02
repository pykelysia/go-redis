package redismq

import "fmt"

func (client *PSubClient) Subscribe(channel string, handler HandlerFunc) error {
	redisClient := client.client
	pubsub := redisClient.Subscribe(client.ctx, channel)
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(client.ctx)
		if err != nil {
			Logger(fmt.Sprintf("ReceiveMessage Error: %v", err))
			continue
		}
		flag := handler(msg.Channel, msg.Payload)
		if !flag {
			break
		}
	}
	return nil
}

func (client *CptClient) Subscribe(channel string, handler HandlerFunc) error {
	redisClient := client.client
	for {
		msg, err := redisClient.BRPop(client.ctx, 0, channel).Result()
		if err != nil {
			Logger(fmt.Sprintf("ReceiveMessage Error: %v", err))
			return err
		}
		flag := handler(msg[0], msg[1])
		if !flag {
			break
		}
	}
	return nil
}
