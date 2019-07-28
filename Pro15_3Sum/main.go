package main

import (
	"fmt"
	"sort"
)

func main()  {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}

/*
将数组按升序排序，从第一个数字开始，从小到大搜索。
然后以下标为k的数字为基础，取3个数字j >k > k。
其中j和i为另外两个数的下标，i，j 分设在数组 [k, len(nums)] 两端
根据 sum 与 0 的大小关系交替向中间逼近
 */
func threeSum(nums []int) [][]int {
	var results [][]int
	target := 0
	// 排序
	sort.Ints(nums)

	for k:=0; k<len(nums)-2; k++ {
		// 若最小的数已经大于0
		if nums[k] > 0 {
			break
		}
		// 若已经处理过相同数据
		if k > 0 && nums[k-1] == nums[k] {
			continue
		}

		// 从[k, len(nums)]的左右边界往中间搜索
		left,right := k+1,len(nums)-1
		for left < right {
			// 搜索到结果
			if nums[k]+nums[left]+nums[right] == target {
				result := []int{nums[k],nums[left],nums[right]}
				results = append(results,result)
			}else if nums[k]+nums[left]+nums[right] > target {
				// 去重
				for right > 0 && nums[right-1] == nums[right] {
					right--
				}
				right--
				continue
			}
			// 去重
			for left < len(nums)-1 && nums[left+1] == nums[left] {
				left++
			}
			left++
		}
	}
	return results
}