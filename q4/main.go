package main

import (
	"fmt"
	"sync"
)

//wg to a pointer so no copies are made and mix everything up 
func Salary(Pay float64, months float64, wg *sync.WaitGroup) {
	salary := Pay * months 
	fmt.Printf("Salary is: $%v \n",salary)
	
	wg.Done()
}

func main() {
	//create the waitgroup 
	var wg sync.WaitGroup

	//the number of goroutines that will be executed 
	wg.Add(2)
	// pass in a pointer to a waitgroup to the function 
	go Salary(3000, 12, &wg)
	go Salary(2000, 12, &wg)


	//wait for all goroutines to be executed
	wg.Wait()

	//once all goroutines are exected this will print
	fmt.Println("Cacluation completed")
}