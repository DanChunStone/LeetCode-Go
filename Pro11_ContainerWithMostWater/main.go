package main

import (
	"fmt"
)

func main()  {
	x := []int{1,8,6,2,5,4,8,3,7}

	fmt.Println(maxArea(x))
}

func maxArea(height []int) int {
	i,j := 0,len(height)-1
	area := 0

	for i < j {
		tmp := (j-i) * Min(height[i],height[j])
		if area < tmp {
			area = tmp
		}

		if height[i] < height[j] {
			i++
		}else {
			j--
		}
	}

	return area
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}