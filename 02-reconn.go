package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Reconn() {
	fmt.Println("Go Redis Reconnect Tutorial")

	client := redis.NewClient(&redis.Options{
		//Addr:     "localhost:6379",
		Addr:     "192.168.61.56:32379",
		Password: "password",
		DB:       0,
	})

	lc := 1
	for {
		val, err := client.Get("k2").Result()
		if err != nil {
			fmt.Printf("%-3d: Error: [%s]\n", lc, err)
		} else {
			fmt.Printf("%-3d: k2=[%s]\n", lc, val)
		}
		lc++

		time.Sleep(1 * time.Second)
	}
}
