package main

import "fmt"

func main()  {
	fmt.Println(letterCombinations("234"))
}

var mapping = map[byte][]byte {
	'2':{'a','b','c'},
	'3':{'d','e','f'},
	'4':{'g','h','i'},
	'5':{'j','k','l'},
	'6':{'m','n','o'},
	'7':{'p','q','r','s'},
	'8':{'t','u','v'},
	'9':{'w','x','y','z'},
}

// 使用动态规划的思想，将前一步的结果记录到
func letterCombinations(digits string) []string {
	var result []string
	if len(digits) == 0 {
		return result
	}

	length := 1
	// 计算结果集大小
	for _,char := range digits {
		if '2' <= char && char <= '9' {
			length = len(mapping[byte(char)])*length
		}
	}

	// 纪录当前产生的结果数
	resultLen := 0
	// 初始化结果集合
	tmpResult := make([][]byte,length)
	for i := 0; i < length; i++ {
		tmpResult[i] = make([]byte,len(digits))
	}

	// 处理第一个字符
	if char := digits[0];'2' <= char && char <= '9' {
		for i := 0; i < len(mapping[char]); i++ {
			tmpResult[i][0] = mapping[char][i]
			resultLen++
		}
	}else {
		tmpResult[0][0] = char
		resultLen++
	}

	for i := 1; i < len(digits); i++ {
		// 如果字符需要转换
		if char := digits[i];'2' <= char && char <= '9' {
			// 复制并产生新结果
			for k := 0; k < len(mapping[char]); k++ {
				for j := 0; j < resultLen; j++ {
					copy(tmpResult[k*resultLen + j][:i],tmpResult[j][:i])
					tmpResult[k*resultLen + j][i] = mapping[char][k]
				}
			}
			resultLen = resultLen*len(mapping[char])
		}else {// 如果字符不需要转换，则直接添加到现有结果的最后一位
			for j := 0; j < resultLen; j++ {
				tmpResult[j][i] = char
			}
		}
	}

	for _,bytes := range tmpResult{
		result = append(result,string(bytes))
	}
	return result
}
