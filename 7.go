package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	sign := 1
	ret := 0
	if x < 0 {
		sign = -1
		x = -x
	}
	for x > 0 {
		ret = ret*10 + (x % 10)
		x /= 10

	}
	ret *= sign
	if ret < math.MinInt32 || math.MaxInt32 < ret {
		return 0
	}
	return ret
}

func main() {
	fmt.Println(reverse(123))        // 321
	fmt.Println(reverse(-123))       // -321
	fmt.Println(reverse(1000000003)) // 0

}
