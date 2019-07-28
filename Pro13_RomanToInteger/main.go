package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
}

func romanToInt(s string) int {
	symbols := map[uint8]int {'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
	result := 0

	for i := 0; i < len(s); {
		// 如果下一个符号比当前符号代表的数字更大，则为形如 4/9
		if i+1 < len(s) && symbols[s[i]] < symbols[s[i+1]] {
			result += symbols[s[i+1]] - symbols[s[i]]
			i += 2
		}else {
			result += symbols[s[i]]
			i++
		}
	}
	return result
}