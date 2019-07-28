package main

import "fmt"

func main()  {
	s := "fsf"
	p := ""
	fmt.Println(isMatch(s,p))
}

func isMatch(s string, p string) bool {
	if len(p) > 0  {
		// 若s已经匹配完毕，而正则p还没有匹配完毕
		if len(s) == 0 {
			// 若正则p长度小于2 或者 不为"_*"的形式
			if len(p) < 2 || p[1] != '*' {
				return false
			}
			return isMatch(s,p[2:])
		// 若s和p均未匹配完
		}else {
			if len(p) > 1 && p[1] == '*' {
				if s[0] == p[0] || p[0] == '.' {
					return isMatch(s[1:],p) || isMatch(s,p[2:])
				}else {
					return isMatch(s,p[2:])
				}
			}else {
				if s[0] == p[0] || p[0] == '.' {
					return isMatch(s[1:],p[1:])
				}else {
					return false
				}
			}
		}
	}

	return len(s) == 0
}