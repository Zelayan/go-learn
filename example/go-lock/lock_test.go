package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/zieckey/etcdsync"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// mutex 锁
func lockTest() {
	var count int64
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			mu.Lock()
			count = count + 1
			mu.Unlock()
		}(&wg)
	}
	wg.Wait()
	print(count)
}

// cas 锁
func atomicTest() {
	var count int64
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 通过
			for {
				old := atomic.LoadInt64(&count)
				if atomic.CompareAndSwapInt64(&count, old, old+1) {
					break
				}
			}

		}(&wg)
	}
	wg.Wait()
}

func redisLockTest() {
	var count int64
	var lockKey = "redisLock"
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				ok, err := rdb.SetNX(context.Background(), lockKey, 1, 5*time.Second).Result()
				if err != nil {
					fmt.Println("Error setting lock:", err)
					return
				}
				if ok {
					count = count + 1
					rdb.Del(context.Background(), lockKey)
					break
				} else {
					// 这里获取不到锁需要增加等待时间防止，redis服务无法正常响应
					time.Sleep(time.Second)
				}

			}
		}(&wg)
	}
	wg.Wait()
	fmt.Println(count)
}

// zkLock
func zkLockTest() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	var count int64
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
				err = l.Lock()
				if err != nil {
					fmt.Println("lock failed:" + err.Error())
					time.Sleep(time.Second)
					continue
				}
				count = count + 1
				l.Unlock()
				return
			}
		}(&wg)
	}
	wg.Wait()
	fmt.Println(count)
}

func etcdLockTest() {
	m, err := etcdsync.New("/lock", 10, []string{"http://127.0.0.1:2379"})
	if m == nil || err != nil {
		log.Printf("etcdsync.New failed")
		return
	}
	m.Lock()
}

func BenchmarkLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lockTest()
	}

}
func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomicTest()
	}
}

func TestRedisLock(t *testing.T) {
	redisLockTest()
}

func TestLockLock(t *testing.T) {
	lockTest()
}

func TestZkLock(t *testing.T) {
	zkLockTest()
}
func TestMain(m *testing.M) {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	m.Run()
}
