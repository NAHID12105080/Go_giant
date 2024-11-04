package main

import (
	"fmt"
)

func sum(a int, b int, c chan int){
	c <- a+b
}




func main() {

result:=make(chan int)
go sum(10, 20, result)
fmt.Println("Result is", <-result)
	
	
}