package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var first int

func f(id int, ch chan bool, once *sync.Once) {
	t := rand.Uint32() % 300
	time.Sleep(time.Duration(t) * time.Millisecond)
	once.Do(func() {
		first = id
	})
	ch <- true
	fmt.Printf("Done %d\n", id)
}

func main() {
	ch := make(chan bool)
	once := sync.Once{}
	for i := 0; i < 5; i++ {
		go f(i, ch, &once)
	}
	for i := 0; i < 5; i++ {
		<-ch
	}
	fmt.Printf("the first is %d\n", first)
}
