package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	MaxGoRoutineCount = 10
)

var (
	result   chan string        = make(chan string, MaxGoRoutineCount)        // 收集结果的chan
	jobQueue chan time.Duration = make(chan time.Duration, MaxGoRoutineCount) // 分发任务的chan
)

// goroutine 超时 demo
func main() {
	// var wg sync.WaitGroup
	// wg.Add(MaxGoRoutineCount)
	// rand.Seed(time.Now().Unix())
	// for i := 0; i < MaxGoRoutineCount; i++ {
	// 	go hello(&wg)
	// }

	// go func() {
	// 	for i := 0; i < 100; i++ { // 有1000个任务要分配出去
	// 		jobQueue <- time.Duration(rand.Intn(10)) * time.Second
	// 	}
	// 	close(jobQueue)
	// }()

	// go func() {
	// 	i := 0
	// 	for v := range result {
	// 		fmt.Println("Result", i, ":", v)
	// 		i++
	// 	}
	// }()

	// wg.Wait()
	// close(result)
	timer()
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	for seconds := range jobQueue { // 死循环, 但是不担心，直到真正有任务之前都卡在这里, 且在关闭workLoad 的时候，所有goroutine 都会离开for 循环
		fmt.Println("job:", seconds.Seconds(), "start")

		select {
		case <-time.After(seconds): // 这个任务要处理n 秒……这里就可能和实际遇到的不一样……意味着真正的任务得再放到某个goroutine 里处理
			fmt.Println("job:", seconds.Seconds(), "success")
			result <- "success!"
		case <-time.After(5 * time.Second): // 设置5秒超时
			result <- "over time!!!"
			fmt.Println("job:", seconds.Seconds(), "timeout")
			continue
		}
	}
}

func timer() {
	//每三秒定时器
	// c := time.Tick(time.Second * 3)
	// for now := range c {
	// 	fmt.Println("now:", now)
	// }

	//三秒超时
	var ch1 = make(chan int, 2)
	//ch1 <- 1
	//for {
	select {
	case <-time.After(time.Second * 3):
		fmt.Println("time out")
		return
	case a := <-ch1:
		time.Sleep(time.Second * 5)
		fmt.Println("running", a)
	default:
		fmt.Println("default.....")
	}
}
