package main

import "fmt"

//by default variables are passed by value
//if you want to pass by reference you can use pointers
//pointers are the memory address of the variable



func chageNumber(number *int){
	*number = 20
}


func main() {
 number := 10
 chageNumber(&number)
 fmt.Println(number)
}