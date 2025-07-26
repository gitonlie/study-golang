package main

import (
	"fmt"
	"time"
)

// 并发Goroutine 1
func goroutinePerfect1() {
	//奇数
	go func() {
		odd := 1
		for {
			if odd > 10 {
				break
			}
			fmt.Println("奇数:", odd)
			odd = odd + 2
		}
	}()

	//偶数
	go func() {
		even := 2
		for {
			if even > 10 {
				break
			}
			fmt.Println("偶数:", even)
			even = even + 2
		}
	}()
}

// GoroutinePerfect2
func goroutinePerfect2(tasks []func()) {
	//任务调度器
	for i := range tasks {
		go func() {
			start := time.Now()
			tasks[i]()
			duration := time.Since(start)
			fmt.Println("任务:", i+1, "耗时:", duration)
		}()
	}
}

func main() {

	fmt.Println("*****************并发Goroutine***************")

	//goroutine1
	goroutinePerfect1()
	time.Sleep(2 * time.Second)

	//goroutine2 调度任务
	job := make([]func(), 0)

	job = append(job, func() {
		time.Sleep(200 * time.Millisecond)
	})

	job = append(job, func() {
		time.Sleep(100 * time.Millisecond)
	})

	job = append(job, func() {
		time.Sleep(300 * time.Millisecond)
	})

	goroutinePerfect2(job)
	time.Sleep(1 * time.Second)

}
