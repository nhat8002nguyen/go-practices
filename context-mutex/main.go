package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Monitor struct {
	ActiveUsers int
	Requests    int
}

func updater(monitor atomic.Value, m *sync.Mutex) {
	for {
		time.Sleep(500 * time.Millisecond)
		m.Lock()
		cur := monitor.Load().(*Monitor)
		cur.ActiveUsers += 300
		cur.Requests += 400
		monitor.Store(cur)
		m.Unlock()
	}
}

func observer(monitor atomic.Value) {
	for {
		time.Sleep(time.Second)
		value := monitor.Load().(*Monitor)
		fmt.Printf("%d %d\n", value.ActiveUsers, value.Requests)
	}
}

func main() {
	monitor := atomic.Value{}
	monitor.Store(&Monitor{0, 0})
	m := sync.Mutex{}

	go updater(monitor, &m)
	go observer(monitor)

	time.Sleep(5 * time.Second)
}
