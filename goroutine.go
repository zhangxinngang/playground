package main

import (
	"fmt"
	"math/rand"
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
		//runtime.Gosched()
		t := time.NewTicker(time.Second * 2)
		<-t.C
		fmt.Println(rand.Perm(5))
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go Deadloop()
	rand.Seed(time.Now().UnixNano())
	//go fmt.Println(1)
	//time.Sleep(time.Second * 10)
	for {
		runtime.Gosched()
	}
}
