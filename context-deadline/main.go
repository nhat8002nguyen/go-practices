package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func accum(id int, count *int32, ctx context.Context) {
	ticker := time.NewTicker(300 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			atomic.AddInt32(count, 1)
		case <-ctx.Done():
			fmt.Printf("Done %d\n", id)
			return
		}
	}
}

func main() {
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	count := int32(0)
	for i := 0; i < 5; i++ {
		go accum(i, &count, ctx)
	}

	<-ctx.Done()
	fmt.Printf("total count: %d\n", count)
}
