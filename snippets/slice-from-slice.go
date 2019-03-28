package main

import (
	"fmt"
)

func main() {
	var s []int = make([]int, 4, 6)
	fmt.Printf("Slice has length %d and capacity %d.\n", len(s), cap(s))
	fmt.Println(s)
	var f = s[0:2]
	fmt.Printf("Slice has length %d and capacity %d.\n", len(f), cap(f))
	fmt.Println(f)

	var e = f[0:3]
	fmt.Printf("Slice has length %d and capacity %d.\n", len(e), cap(e))
	fmt.Println(e)

	s = append(s, 1)
	fmt.Printf("Slice has length %d and capacity %d.\n", len(f), cap(f))
	fmt.Println(f)
	fmt.Printf("Slice has length %d and capacity %d.\n", len(e), cap(e))
	fmt.Println(e)

	var a [5]int
	var b []int = a[0:2]
	fmt.Printf("Slice has length %d and capacity %d.\n", len(b), cap(b))
	fmt.Println(b)
}
