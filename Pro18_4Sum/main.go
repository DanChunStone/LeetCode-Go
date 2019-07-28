package main

import (
	"fmt"
	"sort"
)

func main()  {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums,target))
}

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	if len(nums) < 4 {
		return result
	}

	// 排序
	sort.Ints(nums)
	// 如果target超出数组4数之和的范围
	if min := nums[0]+nums[1]+nums[2]+nums[3]; target < min {
		return result
	}
	max := nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3]+nums[len(nums)-4]
	if target > max {
		return result
	}

	for i := 0; i < len(nums)-3; i++ {
		// 若当前最小的和已经大于target
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		// 若已经处理过相同数据
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i+1; j < len(nums)-2; j++ {
			// 若已经处理过相同数据
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}

			// 从[j, len(nums)]的左右边界往中间搜索
			left,right := j+1,len(nums)-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				// 搜索到结果
				if sum == target {
					tmpResult := []int{nums[i],nums[j],nums[left],nums[right]}
					result = append(result,tmpResult)
				}else if sum > target {
					// 去重
					for right > left && nums[right-1] == nums[right] {
						right--
					}
					right--
					continue
				}
				// 去重
				for left < right && nums[left+1] == nums[left] {
					left++
				}
				left++
			}
		}
	}

	return result
}
