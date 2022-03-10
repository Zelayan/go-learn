package deploy

import (
	"context"
	"github.com/go-redis/redismock/v8"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestDeploy(t *testing.T) {

	Deploy()
}

func TestSercer(t *testing.T) {

	http.HandleFunc("/deploy", func(writer http.ResponseWriter, request *http.Request) {

		err := Deploy()
		if err != nil {

		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("xxxx"))
	})
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		filee, err := ReadFilee(TRACEFILE)
		if err != nil {

		}
		writer.Write([]byte(filee))
	})

	http.ListenAndServe(":1000", nil)
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

func TestName(t *testing.T) {
	content, err := ReadFilee(TRACEFILE)
	if err != nil {

	} else {
		t.Log(content)
	}
}
func TestReadLine(t *testing.T) {
	var current = 5
	line, cont := ReadLine(current, TRACEFILE)
	if line != "" {
		t.Log(line)
	}
	t.Logf("count:%d", cont)
	t.Logf("current:%d", current)
}

func TestCreatfile(t *testing.T) {

	create, err := os.Create(TRACEFILE)
	if err != nil {
		 os.MkdirAll(TRACEFILEDir, os.ModePerm)
		create, _ = os.Create(TRACEFILE)
	}
	create.WriteString("2xxx\n")

}
