package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int] /* ключ */ int /* занечение */) // обращаясь к ключу получаем его значение

	for i, numbers := range nums {
		need := target - numbers

		if idx, ok := m[need]; ok {
			return []int{idx, i}
		}

		m[numbers] = 1
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
