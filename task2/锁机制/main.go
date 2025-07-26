package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter struct {
	num int
	mx  sync.Mutex
}

func (c *Counter) increment() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.num++
}

func (c *Counter) getNum() int {
	return c.num
}

type AtomicCounter struct {
	num atomic.Int64
}

func (c *AtomicCounter) increment() {
	c.num.Add(1)
}

func (c *AtomicCounter) getNum() int64 {
	return c.num.Load()
}
func main() {

	fmt.Println("*****************锁机制***************")
	counter1 := &Counter{}
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10000; i++ {
				counter1.increment()
			}
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("counter1 value:", counter1.getNum())

	//定义一个等待组
	wg := sync.WaitGroup{}
	counter2 := &AtomicCounter{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				counter2.increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("counter2 value:", counter2.getNum())
}
