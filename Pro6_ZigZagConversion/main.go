package main

import "fmt"

func main() {
	s := "PAYPALISHIRING"
	numRows := 3

	fmt.Println(convert(s,numRows))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	// 记录每一行的字符
	str := make([][]byte,numRows)
	index,increase := 0,1

	for _,c := range s{
		// 将当前字符添加到对应行之后
		str[index] = append(str[index],byte(c))
		// 往下或往上移动一行
		index += increase

		// 判断行数增加的方向是否应该改变
		if index == 0 || index == numRows-1 {
			increase = -increase
		}
	}

	var result []byte
	for _,line := range str{
		for _,char := range line {
			result = append(result,byte(char))
		}
	}

	return string(result)
}