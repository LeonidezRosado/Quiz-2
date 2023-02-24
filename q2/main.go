package main

import (
	"fmt"
)

func adding(value1 float64, value2 float64, sum chan float64) {
	//add the value and place into reslut
	result := value1 + value2
	//send the result to the channel sum
	sum <- result
	close(sum)
}

func main() {
	result := make(chan float64)
	go adding(23, 24, result)
	//get data from the channel 
	answer := <- result
	fmt.Println("sum:",answer)
}