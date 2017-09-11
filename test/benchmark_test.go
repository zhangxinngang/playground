package test

import (
	"errors"
	"net/http"
	"testing"
)

// func BenchmarkAA(b *testing.B) {
// 	for i := 0; i < b.N; i++ { //use b.N for looping
// 		Division(4, 5)
// 	}
// }

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}

	return a / b, nil
}

func BenchmarkHttp(b *testing.B) {
	URL := "http://b2bmsa-dev-develop.showbiz.wepiao.com/ads"
	cli := http.Client{}
	for i := 0; i < b.N; i++ {
		cli.Get(URL)
		//fmt.Println(resp)
	}
}

/*
	性能测试：
	go test -test.bench="." -count=5 -cpuprofile=cpu.profile -benchmem
	通过cpu.prof 看性能数据
	go tool pprof popcnt.test cpu.profileq
*/
