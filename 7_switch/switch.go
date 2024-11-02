package main

import (
	"fmt"
)

// import "time"

func main() {
	//simple switch
	// i:=4

	// switch i{
	// case 1:
	// 	fmt.Println("one")	
	
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// case 4:	
	// 	fmt.Println("four")
	// default:
	// 	fmt.Println("default")

		//no need to write break statement
		//default case is optional
		//case can be in any order
		//case can be multiple values
		//case can be multiple expressions
		
		//case can be multiple conditions
		// switch time.Now().Weekday(){
		// case time.Saturday, time.Sunday:
		// 	fmt.Println("It's the weekend")
		// default:
		// 	fmt.Println("It's a weekday")
		// }

		//type switch
			
		whichType := func (i interface{}) {
			switch i.(type){
			case int:
				fmt.Println("It's an integer")
			case string:
				fmt.Println("It's a string")
			default:
				fmt.Println("Unknown type")
			}
		}

		whichType(21)
}
