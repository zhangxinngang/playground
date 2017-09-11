package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	AA int `json:"aa,omitempty"`
	BB int `json:"bb,string,omitempty"`
}

func main() {
	a := A{}
	bytes := []byte("{\"bb\":\"2\"}")
	json.Unmarshal(bytes, &a)
	fmt.Println(a)
	str, _ := json.Marshal(a)
	fmt.Println(string(str))
}
