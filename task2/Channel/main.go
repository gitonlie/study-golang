package main

import (
	"fmt"
	"time"
)

func sendData(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("发送数据>>>[CHANNEL]", i)
	}
	close(ch)
}

func receiveData(ch <-chan int) {
	for v := range ch {
		fmt.Println("[CHANNEL]>>>接收数据", v)
	}
}

// 生产者
func producer(ch chan<- int, data int) {
	ch <- data
	fmt.Println("生产者>>>", data)
}

// 消费者
func consumer(ch <-chan int) {
	for data := range ch {
		fmt.Println(">>>消费者:", data)
	}
}

func main() {

	fmt.Println("*****************Channel***************")

	//Channel1
	ch := make(chan int)
	go sendData(ch)
	go receiveData(ch)

	time.Sleep(3 * time.Second)

	//Channel2
	fmt.Println("生产者>>>CHANNEL>>>消费者")

	channel := make(chan int, 10)

	//启动生产者
	go func() {
		for i := 0; i < 100; i++ {
			producer(channel, i)
		}
	}()

	//启动消费者
	go consumer(channel)

	time.Sleep(5 * time.Second)
}
