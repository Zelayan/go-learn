package locks

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Locks() {
	wg := sync.WaitGroup{}
	var mutex sync.Mutex
	fmt.Println("Locking  (G0)")
	mutex.Lock()
	fmt.Println("locked (G0)")
	wg.Add(3)
	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Locking (G%d)\n", i)
			mutex.Lock()
			fmt.Printf("locked (G%d)\n", i)
			time.Sleep(time.Second * 2)
			mutex.Unlock()
			fmt.Printf("unlocked (G%d)\n", i)
			wg.Done()
		}(i)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("ready unlock (G0)")
	mutex.Unlock()
	fmt.Println("unlocked (G0)")
	wg.Wait()
}

func LockPanic() {
	var mu sync.Mutex
	mu.Lock()
	// mu.Lock() // dead lock
	mu.Unlock()
	// mu.Unlock() err sync unlock of unlucked mutex
}

func RWlocks() {
	var count int
	var wg sync.WaitGroup
	wg.Add(50)
	var rwmu sync.RWMutex

	read := func() {
		rwmu.RLock()
		defer rwmu.RUnlock()
		fmt.Printf("prepare read\n")
		v := count
		fmt.Printf("read value %v\n", v)
		wg.Done()
	}
	write := func() {
		rwmu.Lock()
		defer rwmu.Unlock()
		fmt.Printf("prepare write\n")
		v := rand.Intn(1000)
		count = v
		fmt.Printf("write done %v\n", v)
		wg.Done()
	}
	go func() {
		for i := 0; i < 25; i++ {
			go read()
		}
	}()
	go func() {
		for i := 0; i < 25; i++ {
			go write()
		}
	}()

	wg.Wait()
}
