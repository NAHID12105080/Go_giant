package main

import "fmt"

//there is only for loop in go
func main() {

	//for loop
	// for i:=1; i<=5; i++{
	// 	fmt.Println(i)
	// }

	//for loop in a range , it will print 0 to 3, 4 is exclusive
	for i:=range 4{
		fmt.Println(i)
	}




	//while loop
	// i:=1
	// for i<=5{
	// 	fmt.Println(i)
	// 	i++
	// }

	//infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }

} 