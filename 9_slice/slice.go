package main

import "fmt"

func main() {
	//uninitiazlied slice is nil
	//zero value of slice is nil
	//zeroed values are 0 for numeric types, false for booleans, "" for strings, and nil for interfaces, slices, channels, maps, pointers and functions.

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("slice after append:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)
	
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//2D slice
	fmt.Println("--------2D slice--------")
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoD)

	//slice of slice
	//slice is a reference type
	//slice is a reference to an underlying array
	//slice is a reference to a contiguous segment of an array
	
}