package main

import "fmt"

func main() {
	age := 16
	if age >= 18 {
		fmt.Println("You can vote")
	} else if age >= 16 {
		fmt.Println("You can't vote but you can drive")
	}else {
		fmt.Println("You can't vote")
	}

	//go doesn't have ternary operator, so need to use normal if-else in this type of case
	
}