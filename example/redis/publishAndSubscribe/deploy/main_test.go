package deploy

import (
	"context"
	"github.com/go-redis/redismock/v8"
	"testing"
	"time"
)

func TestDeploy(t *testing.T) {

	Deploy()
}

func TestMockRedis(t *testing.T) {
	db, mock := redismock.NewClientMock()

	mock.ExpectGet("111").SetVal("11111")
	client := TTRedisClient(db)
	result, _ := client.Get(context.TODO(), "111").Result()
	if result != "11111" {
		t.Fatal("err")
	}

}

func TestChannle(t *testing.T) {

	ch1 := make(chan int, 100)

	ch := make(chan int)
	ch2 := make(chan int)

	go func() {
		for range time.Tick(3 * time.Second) {
			ch1 <- 0
		}
	}()

	go func(ch1 chan int) {

		select {
		case <-ch1:
			ch <- 0
		}
	}(ch1)

	go func() {
		for range time.Tick(10 * time.Second) {
			ch2 <- 0
		}
	}()

	select {
	case <-ch:
		println("case1")
		return
	case <-ch2:
		println("case2")
		return
	}
}

func TestTTRedisClient(t *testing.T) {
	if true {
		defer func() {
			t.Log("yes")
		}()
	} else {
		defer func() {
			t.Log("else")
		}()
	}
	defer func() {
		t.Log("xxxxxxzhge ")
	}()

	t.Fatal("xxx")
	t.Log("aaaaa")
}
