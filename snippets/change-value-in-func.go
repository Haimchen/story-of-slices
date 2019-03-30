package main

import (
	"fmt"
)

func write(slice []int) {
	slice[0] = 3
	newSlice := append(slice, 3)
	fmt.Println(newSlice)
}

func main() {
	var slice = make([]int, 4, 5)
	write(slice)
	fmt.Println(slice)
}
