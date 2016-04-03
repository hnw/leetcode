package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	occur := make([]map[rune]bool, len(s)*2)
	repeated := make([]bool, len(s)*2)
	for i, _ := range occur {
		occur[i] = make(map[rune]bool)
	}

	max := 0
	for i, rune := range s {
		for j := 0; j <= i; j++ {
			if repeated[j] {
				continue
			}
			if !occur[j][rune] {
				occur[j][rune] = true
			} else {
				n := len(occur[j])
				if max < n {
					max = n
				}
				occur[j] = nil
				repeated[j] = true
			}
		}
	}

	for i, _ := range repeated {
		if !repeated[i] {
			n := len(occur[i])
			if max < n {
				max = n
			}
		}
	}

	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))          // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))             // 1
	fmt.Println(lengthOfLongestSubstring("abababbxxyzxyzxyz")) // 3
	fmt.Println(lengthOfLongestSubstring("c"))                 // 1
}
