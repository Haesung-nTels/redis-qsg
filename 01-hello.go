package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func Hello() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.61.56:32379",
		Password: "password",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
