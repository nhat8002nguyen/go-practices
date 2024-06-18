package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generator(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- rand.Int()
	}
	close(ch)
	fmt.Println("Generator done")
}

func consumer(id int, ch <-chan int, t time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	for ch := range ch {
		time.Sleep(t)
		fmt.Printf("%d task[%d]\n", id, ch)
	}
	fmt.Printf("consumer %d done\n", id)
}

func main() {
	rand.New(rand.NewSource(42))

	ch := make(chan int, 10)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go generator(ch, &wg)
	go consumer(1, ch, 100*time.Millisecond, &wg)
	go consumer(2, ch, 300*time.Millisecond, &wg)

	wg.Wait()
}
