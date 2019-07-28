package main

import (
	"fmt"
	"sort"
)

func main()  {
	nums := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(nums,1))
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	// 排序
	sort.Ints(nums)

	// 若target不在数组和的范围之内
	min := nums[0] + nums[1] + nums[2]
	if target < min {
		return min
	}
	max := nums[len(nums)-1] + nums[len(nums)-2] + nums[len(nums)-3]
	if target > max {
		return max
	}

	// result从最小和往上搜索
	result := min

	for k:=0; k<len(nums)-2; k++ {
		// 从[k, len(nums)]的左右边界往中间搜索
		left,right := k+1,len(nums)-1
		for left < right {
			// 计算三数之和，并更新result
			sum := nums[k]+nums[left]+nums[right]
			if abs(target-sum) < abs(target-result){
				result = sum
			}

			// 搜索到结果
			if sum == target {
				return sum
			}else if sum < target {// 和小于目标，则第二个数往右移，sum增大
				left++
				continue
			}else {// 和大于目标，则第三个数往左移，sum减小
				right--
			}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}