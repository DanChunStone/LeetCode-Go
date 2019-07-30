package main

import "fmt"

func main()  {
	for _,s := range generateParenthesis(3) {
		fmt.Println(s)
	}
}

func generateParenthesis(n int) []string {
	result := make([]string,0)

	if n == 0 {
		result = append(result,"")
	}else {
		// 对于n组括号组成的合法序列，可以拆解为：
		// 一对括号内包含c对括号合法序列，括号外n-c-1对括号合法序列
		for c := 0; c < n; c++ {
			// 遍历c对括号的所有合法序列
			for _,left := range generateParenthesis(c) {
				// 遍历n-c-1对括号的所有合法序列
				for _,right := range generateParenthesis(n-c-1) {
					result = append(result,"("+left+")"+right)
				}
			}
		}
	}

	return result
}
