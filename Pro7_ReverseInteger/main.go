package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	result := 0

	for x != 0 {
		result = result*10 + x%10
		x = x/10

		if result < math.MinInt32 || result > math.MaxInt32 {
			return 0
		}
	}

	return result
}