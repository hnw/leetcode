package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, v := range nums {
		dict[v] = i
	}
	for i, v := range nums {
		if j, ok := dict[target-v]; ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0 1]
}
