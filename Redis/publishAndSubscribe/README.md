# go-redis publish/subscribe

在数据量较小的情况下，可以使用Redis来实现消息的发布与订阅，来代替Kafka。Kafka对于数据量大的场景下性能卓越，但是对于如此小场景时候，不仅运维成本提升，还用不上多少性能。

不过使用Redis的另一个弊端是消息不能堆积，一旦消费者节点没有消费消息，消息将会丢失。因此需要评估当下场景来选择适合的架构。

此处使用go-redis来实现Redis的发布与订阅。

[官方例子](https://pkg.go.dev/github.com/go-redis/redis/v8#PubSub)
