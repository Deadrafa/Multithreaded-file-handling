package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/Deadrafa/Multithreaded-file-handling/internal/service"
)

func main() {
	cfg := service.ParseFlags()
	rand.Seed(time.Now().UnixNano())

	if err := os.MkdirAll("data", 0755); err != nil {
		panic(err)
	}

	tasks := make(chan int, cfg.Iterations)
	for i := 0; i < cfg.Iterations; i++ {
		tasks <- i
	}
	close(tasks)

	mutexes := make([]sync.Mutex, cfg.NumFiles)
	wg := sync.WaitGroup{}
	wg.Add(cfg.Workers)

	for i := 0; i < cfg.Workers; i++ {
		go service.Worker(i, cfg, tasks, mutexes, &wg)
	}

	wg.Wait()
	fmt.Println("Processing completed")
}
