package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	i := 0
	result := 0
	j := 0
	for j < len(s) {
		i = j
		count := 0
		var a [257]int
		for i < len(s) {
			if a[s[i]-97] == 0 {
				a[s[i]-97]++
				count++
			} else {
				count = 1
			}
			if result < count {
				result = count
			}
			i++
		}
		j++
	}
	return result

}

func main() {
	s := "abced"
	a := lengthOfLongestSubstring(s)
	fmt.Println(a)
}
