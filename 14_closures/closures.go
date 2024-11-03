package main

import "fmt"

func counter () func() int{
	i:=0
	return func() int{
		// this inner function can still update i because it is in the same scope. otherwise i will be out of scope
		i++
		return i
	}
}

func main(){
	c1:=counter()
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())
	

	fmt.Println("-----------first counter completed-----------")
	c2:=counter()
	fmt.Println(c2())
	fmt.Println(c2())
	fmt.Println(c2())
}