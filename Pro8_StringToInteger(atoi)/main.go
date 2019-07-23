package main

import (
	"fmt"
	"unicode"
)

func main()  {
	fmt.Println(myAtoi("-91283472332"))
}

func myAtoi(str string) int {
	MaxInt,MinInt := 2147483647,-2147483648
	isIgnore,isPositive,isNotNum := true,true,false
	result := 0

	for _,char := range str {
		// 遍历到第一个不为‘ ’的字符
		if isIgnore && char != ' ' {
			if char == '-' {
				isPositive = !isPositive
			}else if unicode.IsDigit(char) {
				result += int(char-'0')
			}else if char != '+' {
				return 0
			}
			isIgnore = !isIgnore
		} else if !isIgnore && !isNotNum { // 遍历后续字符
			if unicode.IsDigit(char) {
				result = result*10 + int(char-'0')
			}else {
				isNotNum = true
			}

			// 结果大于IntMax
			if isPositive && result >= MaxInt {
				return MaxInt
			}
			// 结果小于IntMin
			if !isPositive && (-result) <= MinInt {
				return MinInt
			}
		}
	}

	// 若字符串为空或全为空白字符，没有遇到不能忽略的字符，则不转换
	if isIgnore {
		return 0
	}

	if !isPositive {
		return -result
	}
	return result
}