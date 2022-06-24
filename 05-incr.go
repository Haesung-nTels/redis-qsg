package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

// NOT WORK WHEN USE WITH MULTI GO ROUTINE(-m N, N>1)
func Incr() {
	nProcs := flag.Int("p", 1, "number of go routines to run")
	mCnt := flag.Int("m", 100, "number of redis INCR('Count') run")
	clear := flag.Bool("c", false, "Clear redis 'Count'")
	intv := flag.Int("i", 1000, "Report Interval")
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
		fmt.Printf("Go %s Redis Incr %d time(s)\n", myId, *mCnt)

		go func() {
			defer wg.Done()
			client := redis.NewClient(&redis.Options{
				//Addr:     "localhost:6379",
				Addr:     "192.168.61.56:32379",
				Password: "password",
				DB:       0,
			})

			ctx := context.TODO()
			var lval int64 = 0
			failed := 0

			for j := 0; j < *mCnt; {

				val, err := client.Incr(ctx, "Count").Result()
				if err != nil {
					fmt.Printf("%s: %3d: last Count=[%d], Try Again! Error: [%s]\n", myId, j, lval, err)
					failed++
					continue
				} else {
					if failed > 0 {
						fmt.Printf("%s: %3d: Count=[%d]\n", myId, j, val)
						failed = 0
					} else if j%*intv == 0 {
						fmt.Printf("%s: %3d: Count=[%d]\n", myId, j, val)
					}
				}
				j++
				lval = val
			}

		}()
	}

	wg.Wait()

	// summary
	myId := fmt.Sprintf("%d", os.Getpid())
	client := redis.NewClient(&redis.Options{
		//Addr:     "localhost:6379",
		Addr:     "192.168.61.56:32379",
		Password: "password",
		DB:       0,
	})
	val, err := client.Get(ctx, "Count").Result()
	if err != nil {
		fmt.Printf("%s: %3d: Error: [%s]\n", myId, *mCnt, err)
	} else {
		fmt.Printf("%s: %3d: Last Count=[%s]\n", myId, *mCnt, val)
	}
}
