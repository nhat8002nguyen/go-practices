package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func setter(id int, count *int32, ctx context.Context) {
	t := time.NewTicker(300 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Done %d\n", id)
			return
		case <-t.C:
			atomic.AddInt32(count, 1)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	count := int32(0)
	for i := 0; i < 5; i++ {
		go setter(i, &count, ctx)
	}

	time.Sleep(time.Second)

	cancel()
	fmt.Printf("total check: %d\n", count)
	time.Sleep(time.Second)
}
