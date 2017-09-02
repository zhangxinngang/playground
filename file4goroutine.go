package main

import (
	"fmt"
	"sync"
	"time"
)

// 有四个线程1、2、3、4。线程1的功能就是输出1，线程2的功能就是输出2，以此类推.........
// 现在有四个文件ABCD。初始都为空。现要让四个文件呈如下格式：

// A：1 2 3 4 1 2....
// B：2 3 4 1 2 3....
// C：3 4 1 2 3 4....
//D：4 1 2 3 4 1....

var num int

type Printnum struct {
	lock *sync.Mutex
	sum  int
}

func (p *Printnum) Incr() int {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.sum++
	return p.sum
}

var printnum = Printnum{
	lock: new(sync.Mutex),
	sum:  0,
}

func PrintN(i int) {
	go func(i int) {
		for {
			if printnum.Incr() == 100 {
				return
			}
			<-chlista[i-1]
			fmt.Print("a", i)
			cha <- i%4 + 1
		}
	}(i)
	go func(i int) {
		for {
			if printnum.Incr() == 100 {
				return
			}
			<-chlistb[i-1]
			fmt.Print("b", i)
			chb <- i%4 + 1
		}
	}(i)
	go func(i int) {
		for {
			if printnum.Incr() == 100 {
				return
			}
			<-chlistc[i-1]
			fmt.Print("c", i)
			chc <- i%4 + 1
		}
	}(i)
	go func(i int) {
		for {
			if printnum.Incr() == 100 {
				return
			}
			<-chlistd[i-1]
			fmt.Print("d", i)
			chd <- i%4 + 1
		}
	}(i)
}

var ch1a = make(chan int)
var ch2a = make(chan int)
var ch3a = make(chan int)
var ch4a = make(chan int)
var chlista = []chan int{ch1a, ch2a, ch3a, ch4a}

var ch1b = make(chan int)
var ch2b = make(chan int)
var ch3b = make(chan int)
var ch4b = make(chan int)
var chlistb = []chan int{ch1b, ch2b, ch3b, ch4b}

var ch1c = make(chan int)
var ch2c = make(chan int)
var ch3c = make(chan int)
var ch4c = make(chan int)
var chlistc = []chan int{ch1c, ch2c, ch3c, ch4c}

var ch1d = make(chan int)
var ch2d = make(chan int)
var ch3d = make(chan int)
var ch4d = make(chan int)
var chlistd = []chan int{ch1d, ch2d, ch3d, ch4d}

var cha = make(chan int) // 1,2,3,4
var chb = make(chan int) // 1,2,3,4
var chc = make(chan int) // 1,2,3,4
var chd = make(chan int) // 1,2,3,4

func filea() {
	for {
		num := <-cha
		chlista[num-1] <- num
	}
}

func fileb() {
	for {
		num := <-chb
		chlistb[num-1] <- num
	}
}

func filec() {
	for {
		num := <-chc
		chlistc[num-1] <- num
	}
}

func filed() {
	for {
		num := <-chd
		chlistd[num-1] <- num
	}
}

func main() {
	go PrintN(1)
	go PrintN(2)
	go PrintN(3)
	go PrintN(4)

	go filea()
	go fileb()
	go filec()
	go filed()

	chd <- 1
	cha <- 4
	chb <- 3
	chc <- 2

	time.Sleep(time.Second * 1)
}
