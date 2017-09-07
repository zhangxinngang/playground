package main

import "fmt"

//var wg sync.WaitGroup

func main() {
	//wg.Add(-1)
	// cmd := exec.Command("ls", "-l")
	// var out bytes.Buffer
	// cmd.Stdout = &out
	// fmt.Println(cmd.Run())
	// fmt.Println(out.String())

	// c := time.Tick(time.Second * 3)
	// for now := range c {
	// 	fmt.Println("now:", now)
	// }
	// var ch1 = make(chan int, 2)
	// //ch1 <- 1
	// //for {
	// select {
	// case <-time.After(time.Second * 3):
	// 	fmt.Println("time out")
	// 	return
	// case a := <-ch1:
	// 	time.Sleep(time.Second * 5)
	// 	fmt.Println("running", a)
	// }
	//}
	//fmt.Println(removeDuplicates([]int{1, 1, 2, 2}))
	a := new(int)
	b := make([]int, 2)
	fmt.Println(&a, b, a)

	str := "asbdg"
	for _, v := range str {
		fmt.Println(v)
	}
}

func removeDuplicates(nums []int) int {
	mapp := map[int]int{}
	for _, v := range nums {
		mapp[v] = 1
	}
	return len(mapp)
}
