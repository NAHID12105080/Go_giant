package main

import "fmt"


func add(a int, b int) int {
	return a + b
}

func getSubject()(string,string,string){
	return "Math", "Science", "English"
}

func main() {

result := add(1, 2)
fmt.Println(result)

//destructuring
math, science, english := getSubject()	
fmt.Println(math, science, english)


	
}