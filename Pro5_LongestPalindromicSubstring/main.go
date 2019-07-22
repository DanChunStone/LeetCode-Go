package main

import "fmt"

// Given a string s, find the longest palindromic substring in s.
// You may assume that the maximum length of s is 1000.
// 寻找最长回文串
func main() {
	fmt.Println(longestPalindrome("cbbd"))
}


// 动态规划：
// 通过P(i,j)=(P(i+1,j−1) and Si==Sj)
// 初始化一字母和二字母的回文，然后所有的长度（无论奇偶）的回文都可以通过他们长度-2的子串得到
func longestPalindrome(s string) string {
	l := len(s)
	p := [1000][1000] bool{}

	if l == 0 || l == 1 {
		return s
	}

	// 记录子串下标
	aMax,bMax :=0,0

	// 初始化长度为1和长度为2的回文串
	p[l-1][l-1] = true
	for i := 0; i < l - 1; i++ {
		p[i][i] = true
		p[i][i+1] = s[i] == s[i+1]
		if p[i][i+1] {
			aMax,bMax = i,i+1
		}
	}
	
	// 遍历长度由3到L的子串，并且长度由小到大
	// 该循环中，j代表p[a][b]中b-a的大小，即子串长度，易知j的范围为2 - (l-1)
	// 其中，i代表当前循环的下标a，即当前遍历子串的第一个字符，易知i的范围为0-(l-j-1)
	for j := 2; j < l; j++ {
		for i := 0; i < l - j ; i++ {
			p[i][i+j] = p[i+1][i+j-1] && (s[i]==s[i+j])

			// 由于子串长度是递增的，故最后记录的下标代表了最长的回文串
			if p[i][i+j] {
				aMax, bMax = i, i+j
			}
		}
	}

	return s[aMax:bMax+1]
}