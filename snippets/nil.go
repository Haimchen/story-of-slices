package main

import (
	"fmt"
)

func main() {
	var s []int
	fmt.Printf("Slice  has length %d and capacity %d.\n", len(s), cap(s))
	fmt.Println(s)
	var b []int = make([]int, 0, 0)
	fmt.Printf("Slice has length %d and capacity %d.\n", len(b), cap(b))
	fmt.Println(b)
	var t []int = nil
	fmt.Printf("Slice has length %d and capacity %d.\n", len(t), cap(t))
	fmt.Println(t)
}
