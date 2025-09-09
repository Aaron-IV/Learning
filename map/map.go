package main

import "fmt"

func main() {
	weather := map[int]int{
		11: +3,
		12: +6,
		13: +9,
		14: -6,
		15: -12,
	}

	c, ok := weather[30]
	if ok {
		fmt.Println("OKAY")
	} else {
		fmt.Println("BAD")
	}

	fmt.Println(c, ok)
}
