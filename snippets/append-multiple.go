package main

import (
	"fmt"
)

func main() {
	var s []int
	fmt.Printf("Slice has length %d and capacity %d.\n", len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("Slice has length %d and capacity %d.\n", len(s), cap(s))
	s = append(s, 1, 2, 3, 4, 6, 7, 7, 8)
	fmt.Printf("Slice has length %d and capacity %d.\n", len(s), cap(s))
	// how is it 12? Usually it always doubles the capacity!
	// so expected cap values are 2, 4, 8, 16, 32, ...
}
