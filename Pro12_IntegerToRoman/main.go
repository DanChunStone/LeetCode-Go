package main

import "fmt"

func main()  {
	fmt.Println(intToRoman(58))
}

//	Symbol       Value
//	I             1
//	V             5
//	X             10
//	L             50
//	C             100
//	D             500
//	M             1000
func intToRoman(num int) string {
	var result []rune

	// Symbol M
	for num >= 1000 {
		result = append(result,'M')
		num -= 1000
	}
	// Symbol D
	for num >= 500 {
		if num >= 900 {
			result = append(result,'C')
			result = append(result,'M')
			num -= 900
		}else {
			result = append(result,'D')
			num -= 500
		}
	}
	// Symbol C
	for num >= 100 {
		if num >= 400 {
			result = append(result,'C')
			result = append(result,'D')
			num -= 400
		}else {
			result = append(result,'C')
			num -= 100
		}
	}
	// Symbol L
	for num >= 50 {
		if num >= 90 {
			result = append(result,'X')
			result = append(result,'C')
			num -= 90
		}else {
			result = append(result,'L')
			num -= 50
		}
	}
	// Symbol X
	for num >= 10 {
		if num >= 40 {
			result = append(result,'X')
			result = append(result,'L')
			num -= 40
		}else {
			result = append(result,'X')
			num -= 10
		}
	}
	// Symbol V
	for num >= 5 {
		if num >= 9 {
			result = append(result,'I')
			result = append(result,'X')
			num -= 9
		}else {
			result = append(result,'V')
			num -= 5
		}
	}
	// Symbol I
	for num >= 1 {
		if num >= 4 {
			result = append(result,'I')
			result = append(result,'V')
			num -= 4
		}else {
			result = append(result,'I')
			num -= 1
		}
	}

	return string(result)
}