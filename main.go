package main

import (
	"fmt"
	"redismq"
	"strconv"
	"sync"
	"time"
)

func main() {
	enum := 0
	fmt.Scan(&enum)
	switch enum {
	case 1:
		testPSubClient()
	case 2:
		testCptClient()
	case 3:
		testCptClientMore()
	}
}

func testPSubClient() {
	wg := sync.WaitGroup{}
	client := redismq.DefultClient()

	wg.Add(1)
	go func() {
		fmt.Println("Start Subscribe:")

		_ = client.Subscribe("channel1", func(channel string, message string) bool {
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

func testCptClient() {
	wg := sync.WaitGroup{}
	client := redismq.DefultCptClient()

	client.Publish("channel", "task1")
	client.Publish("channel", "task2")
	client.Publish("channel", "task3")

	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go func() {
			fmt.Printf("channel %d Start:\n", i)

			client.Subscribe("channel", func(channel string, message string) bool {
				fmt.Println(message)
				return false
			})

			fmt.Printf("channel %d Over.\n", i)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Over.")
}

func testCptClientMore() {
	wg := sync.WaitGroup{}
	client := redismq.DefultCptClient()

	client.Publish("channel", "task1")
	client.Publish("channel", "task2")
	client.Publish("channel", "task3")
	client.Publish("channel", "task4")
	client.Publish("channel", "task5")
	client.Publish("channel", "task6")
	client.Publish("channel", "task7")

	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go func() {
			fmt.Printf("channel %d Start:\n", i)

			client.Subscribe("channel", func(channel string, message string) bool {
				fmt.Println("channel " + strconv.Itoa(i) + ": " + message)
				return true
			})

			fmt.Printf("channel %d Over.\n", i)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Over.")
}
