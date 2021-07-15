package main

import (
	"fmt"
)

func fibb(n int) int{
	if n <= 1 {
		return n
	} else{
		return fibb(n-1) + fibb(n-2)
	}
}

func main() {
	//For a fixed number of iterations
	numbIterations := 10
	
	//For user definded number of iterations
	//var numbIterations int
	//fmt.Scan(&numbIterations)
	
	//Validation
	//if numbIterations <= 0 {
	//	fmt.Println("Error: input must be a pos int")
	//}
	
	
	for i := 0; i < numbIterations; i++ {
		fmt.Println(fibb(i))
	} 
}
