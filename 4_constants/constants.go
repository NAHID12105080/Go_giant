package main

import "fmt"

//short hand declaration can't be used outside for variable
//var a:=8
//but we can use it inside function

func main() {
	const a int = 8

	// inside short hand declaration we can't use const
	//const b:=3

	fmt.Println(a)

	const (
		port = 8080
		host = "localhost"

	)

	fmt.Println(port, host)
	 

}