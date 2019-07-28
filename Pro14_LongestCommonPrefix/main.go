package main

import "fmt"

func main()  {
	strs := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串为基础，对比每一个字符串的 i位
	for i := 0; i < len(strs[0]); i++ {
		// 从第二个字符串开始对比
		for j := 1; j < len(strs); j++ {
			// 若i位超出了当前字符串的长度 或 该位的字符不匹配
			if i >= len(strs[j]) || strs[0][i] != strs[j][i]  {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}