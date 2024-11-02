package main

import "fmt"

func main() {
	name:=[]int {1,2,3,4,5,6}
	for i, v := range name {
		fmt.Println(i, v)
		fmt.Println("----------------")
	}

	// for i := range name {
	// 	fmt.Println(i)
	// 	fmt.Println("----------------")

	// }

	for _, v := range name {
		fmt.Println(v)
		fmt.Println("----------------")

	}

	// for i := 0; i < len(name); i++ {
	// 	fmt.Println(name[i])
	// }
	
}