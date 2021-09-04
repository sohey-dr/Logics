package main

import (
        "github.com/sonots/go-resque" 
        _ "github.com/sonots/go-resque/godis"
        "github.com/simonz05/godis/redis"
)

func main() {
        client := redis.New("tcp:127.0.0.1:6379", 0, "") // Godis redis client
        enqueuer := resque.NewRedisEnqueuer("godis", client)

        enqueuer.Enqueue("default", "MyClass", 1, 2, "woot")
}