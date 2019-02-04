package main

import (
	"fmt"
	"time"
)

var counter = 0

func main() {
	for i := 0; i < 900; i++ {
		incr()
	}
	time.Sleep(time.Millisecond * 10)
}
func incr() {
	counter++
	fmt.Println(counter)
}
