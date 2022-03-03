package deploy

import (
	"context"
	"github.com/go-redis/redismock/v8"
	"testing"
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