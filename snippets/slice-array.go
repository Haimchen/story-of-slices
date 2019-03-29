package main

import (
	"fmt"
)

func main() {
	var arr [20]int
	for idx, _ := range arr {
		s := arr[idx : idx+1]
		fmt.Printf("Slice No %d: Len=%d, Cap=%d\n", idx+1, len(s), cap(s))
	}
}
