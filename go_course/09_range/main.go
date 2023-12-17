package main

import "fmt"

func main() {
	ids := []int{33, 14, 15, 53, 24, 64}
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
}
