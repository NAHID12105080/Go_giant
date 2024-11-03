package main

import "fmt"

func main(){
	type person struct {
		name string
		age int
	}

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	// & prefix yields a pointer to the struct
	//outputs the memory address of the struct person
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)


	//slices ary passed by reference

}