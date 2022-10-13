package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		fmt.Println("another goroutin")
	}()
	fmt.Println("STOP")
	<-ctx.Done()
	fmt.Println("then starts again..,")
}
