package main

import (
	"fmt"
	"redismq"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	client := redismq.DefultClient()

	wg.Add(1)
	go func() {
		fmt.Println("Start Subscribe:")

		_ = client.Subscribe("channel1", func(channel string, message ...string) bool {
			fmt.Println(message)
			return false
		})

		fmt.Println("Subscribe Over.")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("Start Publish:")
		err := client.Publish("channel1", "hello"+" my redis mq")
		if err != nil {
			fmt.Println("Publish Error:", err)
			wg.Done()
			return
		}
		fmt.Println("Publish Over.")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Over.")
}
