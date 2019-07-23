package main

import "fmt"

func main()  {
	fmt.Println(isPalindrome(120))
}

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x % 10 == 0 {
		return false
	}

	// 记录x后半部分的反转
	reverseNum := 0
	// 将x的前半部分保留在x，后半部分保存在reverse。
	// 并且当x位数为奇数时，多余的一位也保留在reverse
	for x > reverseNum {
		reverseNum = reverseNum*10 + x%10
		x = x / 10
	}

	return x == reverseNum || x == reverseNum/10
}