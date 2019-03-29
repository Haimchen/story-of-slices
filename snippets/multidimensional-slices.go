package main

import (
	"fmt"
	"time"
)

func trackTime(start time.Time, name string) {
	elapsedTime := time.Since(start)
	fmt.Printf("function \"%s\" took %s.\n", name, elapsedTime)
}

func iterateRowFirst(slice [][]int) {
	defer trackTime(time.Now(), "rowFirst")
	var count int
	for _, val := range slice {
		for i, _ := range val {
			val[i] = count
			count++
		}
	}
}

// we need a symmetrical multidimensional array for this to work
func iterateColFirst(slice [][]int) {
	defer trackTime(time.Now(), "colFirst")
	var count int
	for i, _ := range slice[0] {
		for n := 0; n < len(slice); n++ {
			slice[n][i] = count
			count++
		}
	}
}

func main() {
	var n = 10000
	var slice = make([][]int, n)
	for idx, _ := range slice {
		slice[idx] = make([]int, n)
	}
	iterateRowFirst(slice)
	// fmt.Println(slice)
	iterateColFirst(slice)
	// fmt.Println(slice)
}
