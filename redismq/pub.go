package redismq

func (client *Client) Publish(channel string, message interface{}) error {
	redisClient := client.client
	err := redisClient.Publish(client.ctx, channel, message).Err()
	return err
}
