package redismq

func (client *PSubClient) Publish(channel string, message interface{}) error {
	redisClient := client.client
	err := redisClient.Publish(client.ctx, channel, message).Err()
	return err
}

func (client *CptClient) Publish(channel string, message interface{}) error {
	redisClient := client.client
	redisClient.LPush(client.ctx, channel, message)
	return nil
}
