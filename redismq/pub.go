package redismq

func (client *PSubClient) Publish(channel string, message interface{}) error {
	redisClient := client.client
	err := redisClient.Publish(client.ctx, channel, message).Err()
	return err
}
