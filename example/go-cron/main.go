package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	// 定义一个cron，开启秒级别执行
	c := cron.New(cron.WithSeconds())
	cronPool := make(map[string]cron.EntryID)
	// 每5秒执行一次
	every5s, err := c.AddFunc("@every 5s", func() {
		time.Now().Format("2006-01-02 15:04:05")
		fmt.Println("@every 5s")
	})
	cronPool["every5s"] = every5s

	// 每10s执行一次
	cron10s, err := c.AddFunc("@every 10s", func() {
		time.Now().Format("2006-01-02 15:04:05")
		fmt.Println("@every 10s")
	})
	cronPool["cron10s"] = cron10s
	if err != nil {
		panic(err)
	}

	go func() {
		// 20s后取消 cron10s
		time.Sleep(time.Second * 20)
		id := cronPool["cron10s"]
		c.Remove(id)
	}()

	// 开启定时任务池
	c.Run()

}
