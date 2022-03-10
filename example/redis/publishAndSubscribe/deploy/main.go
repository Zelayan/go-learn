package deploy

import (
	"bufio"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func Deploy() error {

	deployDone := make(chan struct{})
	runShDone := make(chan struct{}, 1)

	go subscribe(deployDone)

	time.Sleep(5 * time.Second)
	runShDone <- struct{}{}

	ticker := time.NewTicker(10 * time.Second)

	select {
	case <-deployDone:
		fmt.Println("redis")
		return nil
	case <-runShDone:
		fmt.Println("deploy done")
		return nil
	case <- ticker.C:
		fmt.Println("ticker")
		return nil
	}

	return nil

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

const TRACEFILE = "/Users/chenze/tmp/1"
const TRACEFILEDir = "/Users/chenze/tmp"

func GetDeployProcess() error{
	file, err := os.Open(TRACEFILE)
	if err != nil {
		return fmt.Errorf("not in deploy:%w", err)
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				time.Sleep(100 * time.Millisecond)
			} else {
				break
			}
		}
		fmt.Print(string(line))
	}
	return nil
}
func ReadFilee(filepath string) (string, error){
	content ,err :=ioutil.ReadFile(filepath)
	if err !=nil {
		return "", err
	}
	return string(content), nil
}

func ReadFile(path string) (bool, error){
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func ReadLine(lineNumber int, path string) (string, int) {
	file, _ := os.Open(path)
	fileScanner := bufio.NewScanner(file)
	lineCount := 1

	result := ""
	for fileScanner.Scan(){
		if lineCount == lineNumber{
			result =  fileScanner.Text()
		}
		lineCount++

	}
	defer file.Close()
	return result, lineCount - 1
}

