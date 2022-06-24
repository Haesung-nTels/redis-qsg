package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

func MultiIncr() {
	nProcs := flag.Int("p", 1, "number of go routines to run")
	mCnt := flag.Int("m", 100, "number of redis INCR('Count') run")
	clear := flag.Bool("c", false, "Clear redis 'Count'")
	intv := flag.Int("i", 100, "Report Interval")
	flag.Parse()

	if *clear {
		client := redis.NewClient(&redis.Options{
			//Addr:     "localhost:6379",
			Addr:     "192.168.61.56:32379",
			Password: "password",
			DB:       0,
		})

		ctx := context.TODO()
		client.Del(ctx, "Count")

		return
	}

	var wg sync.WaitGroup

	for i := 0; i < *nProcs; i++ {
		wg.Add(1)

		myId := fmt.Sprintf("%d.%d", os.Getpid(), i)
		fmt.Printf("Go %s Redis Multi Incr %d time(s)\n", myId, *mCnt)

		go func() {
			defer wg.Done()
			client := redis.NewClient(&redis.Options{
				//Addr:     "localhost:6379",
				Addr:     "192.168.61.56:32379",
				Password: "password",
				DB:       0,
			})

			ctx := context.TODO()
			pipe := client.Pipeline()
			failed := 0

			for j := 0; j < *mCnt; {
				// val, err := client.Incr(ctx, "Count").Result()
				incr := pipe.Incr(ctx, "Count")
				if _, err := pipe.Exec(ctx); err != nil {
					if failed == 0 {
						j--
					}
					fmt.Printf("%s: %3d: Try Again! Pipe Error: [%s]\n", myId, j, err)
					failed++
					time.Sleep(1 * time.Second)
					continue
				} else {
					if failed > 0 {
						fmt.Printf("%s: %3d: Count=[%d]\n", myId, j, incr.Val())
						failed = 0
					} else {
						if j%*intv == 0 {
							fmt.Printf("%s: %3d: Count=[%d]\n", myId, j, incr.Val())
						}
					}
				}
				j++
			}

			val, err := client.Get(ctx, "Count").Result()
			if err != nil {
				fmt.Printf("%s: %3d: Error: [%s]\n", myId, *mCnt, err)
			} else {
				fmt.Printf("%s: %3d: Last Count=[%s]\n", myId, *mCnt, val)
			}
		}()
	}

	wg.Wait()
}
