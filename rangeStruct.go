package main

import (
	"fmt"
	"reflect"
)

type Info struct {
	Name string
	Age  int
}

func main() {
	a := Info{Name: "aaa", Age: 11}
	types := reflect.TypeOf(a)
	values := reflect.ValueOf(a)
	for i := 0; i < types.NumField(); i++ {
		fmt.Println(types.Field(i).Name, values.Field(i).Interface())
	}
}
