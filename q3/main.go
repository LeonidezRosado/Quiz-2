package main

import (
	"fmt"
	"time"
)

//creat function named user 
//print out the string that is passed in 3 times 
//use the sleep function here to delay for 100 milliseconds
func user(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	//by writing go infront of the function call it is now a goroutine
	go user("first")
	 user("second")

	 //use time function so goroutine can run 
	time.Sleep(time.Second)
}