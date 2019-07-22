package main

import "fmt"

/**
给定一个字符串，请你找出其中不含有重复字符的 最长子序列 的长度。
 */
func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	length := make([]int,len(s))
	max := 1

	length[0] = 1
	// 遍历每个字符，判断当前位置 i 的字符往前能构成的最长子序列长度
	// length[i] = 从i位置往前搜索，直到找到相同的字符 或 到达i-1位置的最大length
	for i := 1; i < len(s); i++ {
		length[i] = 1
		// 往前搜索，直到找到相同的字符 || 到达i-1位置的最大length
		for j := 1; j <= length[i-1]; j++ {
			if s[i-j] == s[i] {
				break
			}
			length[i]++
		}

		// 记录最大长度
		if length[i] > max {
			max = length[i]
		}
	}

	return max
}