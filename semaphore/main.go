package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

type Task struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	var t Task
	sema := semaphore.NewWeighted(int64(runtime.NumCPU()))
	errGroup, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < 100; i++ {
		fmt.Println(runtime.NumGoroutine())
		// TryAcquire returns a bool and not block, this Acquire blocks
		if err := sema.Acquire(ctx, 1); err != nil {
			log.Fatal(err)
		}
		i := i
		errGroup.Go(func() error {
			defer sema.Release(1)
			res, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", i))
			if err != nil {
				return err
			}
			defer res.Body.Close()
			if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
				return err
			}
			fmt.Println(t.Title)
			return nil
		})
	}
	if err := errGroup.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("there were %d cpus in this host", runtime.NumCPU())
}
