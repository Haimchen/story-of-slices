package main

import (
	"fmt"
)

func dropLast(arr [3]int) [2]int {
	var arr2 [2]int
	arr2[0] = arr[0]
	arr2[1] = arr[1]
	return arr2

}

func dropLastSlice(slice []int) []int {
	end := len(slice)
	slice2 := slice[0:end]
	return slice2
}

func main() {
	var arr = [3]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println(&arr[0])

	arr2 := dropLast(arr)
	fmt.Println(arr2)
	fmt.Println(&arr2[0])
	// new memory location

	var slice = make([]int, 3)
	fmt.Println(slice)
	fmt.Println(&slice[0])
	slice2 := dropLastSlice(slice)
	fmt.Println(slice2)
	fmt.Println(&slice2[0])
	// same memory location

}
