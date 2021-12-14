# Go常见面试题
## 
## 并发编程
### goroutine泄漏
协程泄露是指协程创建后，长时间得不到释放，并且还在不断地创建新的协程，最终导致内存耗尽，程序崩溃。常见的导致协程泄露的场景有以下几种：
1. 缺少接收器，导致发送阻塞
    
    ```go
    func query() int {
	ch := make(chan int)
	for i := 0; i < 1000; i++ {
		go func() { ch <- 0 }()
	}
	return <-ch
    }

    func main() {
        for i := 0; i < 4; i++ {
            query()
            fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
        }
    }
    // goroutines: 1001
    // goroutines: 2000
    // goroutines: 2999
    // goroutines: 3998
    ```
2. 缺少发送器，导致接受阻塞

    
3. 死锁(dead lock)

    两个或两个以上的协程在执行过程中，由于竞争资源或者由于彼此通信而造成阻塞，这种情况下，也会导致协程被阻塞，不能退出。

4. 无限循环

    这个例子中，为了避免网络等问题，采用了无限重试的方式，发送 HTTP 请求，直到获取到数据。那如果 HTTP 服务宕机，永远不可达，导致协程不能退出，发生泄漏。

    ```go   
    func request(url string, wg *sync.WaitGroup) {
        i := 0
        for {
            if _, err := http.Get(url); err == nil {
                // write to db
                break
            }
            i++
            if i >= 3 {
                break
            }
            time.Sleep(time.Second)
        }
        wg.Done()
    }

    func main() {
        var wg sync.WaitGroup
        for i := 0; i < 1000; i++ {
            wg.Add(1)
            go request(fmt.Sprintf("https://127.0.0.1:8080/%d", i), &wg)
        }
        wg.Wait()
    }
    ```
