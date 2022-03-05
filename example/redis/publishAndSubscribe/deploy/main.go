package deploy

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func Deploy() error {

	deployDone := make(chan struct{})
	runShDone := make(chan struct{})

	go subscribe(deployDone)
	time.Sleep(5 * time.Second)
	runShDone <- struct{}{}

	select {
	case <-deployDone:
		fmt.Println("redis")
		return nil
	case <-runShDone:
		fmt.Println("deploy done")
		return nil
	}

}

func NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	todo := context.TODO()
	_, err := client.Ping(todo).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func TTRedisClient(client *redis.Client) *redis.Client {
	return client
}

func subscribe(ch chan struct{}) {

	client, _ := NewRedisClient()
	defer client.Close()
	sub := client.Subscribe(context.TODO(), "deploy")
	_, err := sub.ReceiveMessage(context.TODO())

	if err != nil {
		fmt.Println(err.Error())
	}
	channel := sub.Channel()
	for m := range channel {
		fmt.Println(m)
		ch <- struct{}{}
	}
}
