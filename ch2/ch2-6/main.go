package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	task := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case i := <-task:
				fmt.Println("get ", i)
			default:
				fmt.Println("have not been cancelled")
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()
	time.Sleep(3 * time.Second)
	for i := 0; 5 > i; i++ {
		task <- i
	}
	cancel()
}
