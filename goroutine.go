package main

import (
	"fmt"
	"runtime"
	"time"
)

func Deadloop() {
	for {
		// select {
		// case <-time.After(time.Second * 2):
		// 	//runtime.Gosched()
		// 	fmt.Println(22)
		// }
		runtime.Gosched()
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go Deadloop()
	go fmt.Println(1)
	time.Sleep(time.Second * 10)
}
