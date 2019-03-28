package main

import (
	"fmt"
)

func main() {
	i := make([]int, 5, 5)
	fmt.Printf("Slice has length %d and capacity %d.\n", len(i), cap(i))
	fmt.Println(i)

	s := make([]int, 5, 7)
	fmt.Printf("Slice has length %d and capacity %d.\n", len(s), cap(s))
	fmt.Println(s)

	k := make([]int, 3, 5)
	fmt.Printf("Slice has length %d and capacity %d.\n", len(k), cap(k))
	fmt.Println(k)
}
