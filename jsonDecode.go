package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	AA *int `json:"aa,omitempty"`
	BB *int `json:"bb,omitempty"`
}

func main() {
	a := A{}
	bytes := []byte("{\"bb\":2}")
	json.Unmarshal(bytes, &a)
	fmt.Printf("%#v\n", a)
	fmt.Println(*a.BB)
	str, _ := json.Marshal(a)
	fmt.Println(string(str))
}
