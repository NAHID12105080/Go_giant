package main

import "fmt"

func modifySlice(s []int) {
    s[0] = 99 // This will modify the original slice
}

func main() {
    nums := []int{1, 2, 3}
    fmt.Println("Before:", nums)
    modifySlice(nums) // Pass the slice to the function
    fmt.Println("After:", nums) // nums[0] is now 99
}
