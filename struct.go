package main

import (
	"fmt"
)

type A interface {
	SetNames(n string)
}

type Info struct {
	Name string
	Age  int
}

func (i Info) SetNames(n string) {

}

func (i *Info) SetName(n string) {
	i.Name = n
}

func main() {
	i := Info{}
	var inter A
	inter = i
	_, ok := interface{}(i).(A)
	fmt.Println(ok, inter)

	var _ A = Info{}
}
