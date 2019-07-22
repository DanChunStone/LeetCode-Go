package main

import "fmt"

/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
You may assume nums1 and nums2 cannot be both empty.
 */
func main() {
	nums1 := []int{1,3}
	nums2 := []int {2}
	fmt.Printf("%f\n",findMedianSortedArrays(nums1,nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 计算数组总长度
	length := len(nums1)+len(nums2)
	n := length / 2 + length % 2

	// 访问数组1和2的指针
	i,j := 0,0
	var index int

	for n > 0 {
		// 数组2已空 或者 数组1当前元素更小，且未遍历完
		if j >= len(nums2) || i < len(nums1) && nums1[i] <= nums2[j] {
			index = 1
			i++
		}else {
			index = 2
			j++
		}
		n--
	}

	var result float64;

	// 获取中间位置的数
	if index == 1 {
		result = float64(nums1[i-1])
	}else {
		result = float64(nums2[j-1])
	}

	// 若数组总长度为偶数，则还要计算下一个数的平均数
	if length%2 == 0 {
		if j >= len(nums2) || i < len(nums1) && nums1[i] <= nums2[j] {
			result = (result + float64(nums1[i])) / 2
		}else {
			result = (result + float64(nums2[j])) / 2
		}
	}
	// 若为奇数，则直接返回中间位置的数
	return result
}
