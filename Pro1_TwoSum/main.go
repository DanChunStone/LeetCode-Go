package main

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
 */
func main() {
	
}

// case 1
// func twoSum(nums []int, target int) []int {
// 	for i := range nums {
// 		for j := i + 1; j < len(nums); j++ {
// 			if target == nums[i]+nums[j] {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }


// case 2
func twoSum(nums []int, target int) []int {
	numbers := make(map[int]int)
	for i := range nums {
		complement := target - nums[i]
		index, ok := numbers[complement]
		if ok {
			return []int{index, i}
		}
		numbers[nums[i]] = i
	}
	return nil
}