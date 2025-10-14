package main

import (
	"fmt"
)

func main() {
	arr := [5]int{5, 66, 7, 100, 1}

	fmt.Println("Нулевой элемент;", arr[0])

	arr[0] *= 5

	fmt.Println(arr)
}
