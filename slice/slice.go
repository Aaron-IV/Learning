package main

import "fmt"

func main() {
	intslice := make([]int, 3)

	fmt.Println((intslice), len(intslice), cap(intslice))
}
