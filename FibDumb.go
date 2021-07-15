package main

import (
	"fmt"
)

func main(){
    x := 0
    y := 1
    f := 0
    r := 10 //change to input later
    //run for 10 iterations
    fmt.Println(f)
    for i := 0; i < r; i++ {
      f := x + y
      fmt.Println(f)
      x = y
      y = f
    }
 }
 
 //finds the first 10 terms of fibb using an interative (boring) approach
