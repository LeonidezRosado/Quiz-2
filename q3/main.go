package main

import (
	"fmt"
	"time"
)

func user(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go user("first")
	 user("second")

	time.Sleep(time.Second)
}