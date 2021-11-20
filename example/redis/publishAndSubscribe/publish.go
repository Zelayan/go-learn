// publish msg
package publishAndSubscribe

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
    "log"
    "time"
)

func pubMessage(channel, msg string) {
    rdb := redisConnect()
    rdb.Publish(context.Background(), channel, msg)
}

func redisConnect() (rdb *redis.Client) {
    rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    return
}

func Publish(channel string) {
    mgsList := []string{"hello", "world"}

    for _, msg := range mgsList {
        pubMessage(channel, msg)
        log.Printf("send %v to %s\n", msg, channel)
    }
}

func subMessage(channel string) {
    rdb := redisConnect()
    subscribe := rdb.Subscribe(context.Background(), channel)
    _, err := subscribe.ReceiveMessage(context.Background())
    if err != nil {
        log.Println("receive message failed", err)
    }

    ch := subscribe.Channel()

    // Consume messages.
    for msg := range ch {
        fmt.Println(msg.Channel, msg.Payload)
    }
}

func Subscribe(channel string) {
    subMessage(channel)
}

func OfficalCase() {
    ctx := context.TODO()
    rdb := redisConnect()
    pubsub := rdb.Subscribe(ctx, "mychannel1")

    // Wait for confirmation that subscription is created before publishing anything.
    _, err := pubsub.Receive(ctx)
    if err != nil {
        panic(err)
    }

    // Go channel which receives messages.
    ch := pubsub.Channel()

    // Publish a message.
    err = rdb.Publish(ctx, "mychannel1", "hello").Err()
    if err != nil {
        panic(err)
    }

    time.AfterFunc(time.Second, func() {
        // When pubsub is closed channel is closed too.
        _ = pubsub.Close()
    })

    // Consume messages.
    for msg := range ch {
        fmt.Println(msg.Channel, msg.Payload)
    }
}