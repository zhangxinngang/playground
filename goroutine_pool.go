package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type GoroutinePool struct {
	Channel chan func()
	Num     int
}

func (g *GoroutinePool) Init(num int) {
	g.Channel = make(chan func(), num)
	g.Num = num
}

func (g *GoroutinePool) Add(f func()) {
	g.Channel <- f
}

func (g *GoroutinePool) Run() {
	for i := 0; i < g.Num; i++ {
		go func() {
			for {
				task := <-g.Channel
				task()
			}
		}()
	}
}

func print() {
	fmt.Println("A")
}

func main() {
	pool := GoroutinePool{}
	pool.Init(3)
	go func() {
		for i := 0; i < 10; i++ {
			pool.Add(print)
			fmt.Println("asdgs")
		}
	}()
	pool.Run()
	time.Sleep(time.Second * 2)
}
