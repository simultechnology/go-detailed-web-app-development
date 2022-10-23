package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("037 start!")

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("i: %d\n", i)
			wg.Done()
		}()
	}
	wg.Wait()
	for j := 0; j < 5; j++ {
		wg.Add(1)
		go func(x int) {
			fmt.Printf("j: %d\n", x)
			wg.Done()
		}(j)
	}
	wg.Wait()
	for k := 0; k < 5; k++ {
		time.Sleep(time.Millisecond * 10)
		k := k
		wg.Add(1)
		go func() {
			fmt.Printf("k: %d\n", k)
			wg.Done()
		}()
	}
	wg.Wait()
}
