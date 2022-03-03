package deploy

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func Deploy() error {

	deployDone := make(chan struct{}, 1)
	runShDone := make(chan struct{}, 1)

	go subscribe(deployDone)
	go runShDoneFun(runShDone)

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
	sub := client.Subscribe(context.TODO(), "test")
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

func runShDoneFun(ch chan struct{}) {
	time.Sleep(10 * time.Second)
	ch <- struct{}{}
}
