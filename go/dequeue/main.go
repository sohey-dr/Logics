package main

import (
    "fmt"

    "github.com/go-redis/redis"
)

func main() {
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    fmt.Println("Redis client:", client)

}

func StringGetSet(client *redis.Client) {
    key := "StringGetSet_Key"
    // Set
    err := client.Set(key, "StringGetSet_Val", 0).Err()
    if err != nil {
        fmt.Println("redis.Client.Set Error:", err)
    }

    // Get
    val, err := client.Get(key).Result()
    if err != nil {
        fmt.Println("redis.Client.Get Error:", err)
    }
    fmt.Println(val)
}

func ListGetSet(client *redis.Client) {
    key := "ListGetSet_Key"
    // Set
    listVal := []string{"val1", "va2", "val3"}
    err := client.RPush(key, listVal).Err()
    if err != nil {
        fmt.Println("redis.Client.RPush Error:", err)
    }

    // Get
    // Get by lrange
    lrangeVal, err := client.LRange(key, 0, -1).Result()
    if err != nil {
        fmt.Println("redis.Client.LRange Error:", err)
    }
    fmt.Println(lrangeVal)
    // Get by lindex
    lindexVal, err := client.LIndex(key, 2).Result()
    if err != nil {
        fmt.Println("redis.Client.LIndex Error:", err)
    }
    fmt.Println(lindexVal)
}

func HashGetSet(client *redis.Client) {
    key := "HashGetSet_Key"
    // Set
    for field, val := range map[string]string{"field1": "val1", "field2": "val2"} {
        fmt.Println("Inserting", "field:", field, "val:", val)
        err := client.HSet(key, field, val).Err()
        if err != nil {
            fmt.Println("redis.Client.HSet Error:", err)
        }
    }

    // Get
    // HGet(key, field string) *StringCmd
    hgetVal, err := client.HGet(key, "field1").Result()
    if err != nil {
        fmt.Println("redis.Client.HGet Error:", err)
    }
    fmt.Println(hgetVal)

    // HGetAll
    hgetallVal, err := client.HGetAll(key).Result()
    if err != nil {
        fmt.Println("redis.Client.HGetAll Error:", err)
    }
    // fmt.Println("reflect.TypeOf(hgetallVal):", reflect.TypeOf(hgetallVal)) // map[string]string
    fmt.Println(hgetallVal)
}