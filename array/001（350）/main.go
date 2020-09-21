package main

import "fmt"

//第350题：两个数组的交集
//给定两个数组，编写一个函数来计算它们的交集。
//示例 1:
//
//输入: nums1 = [1,2,2,1], nums2 = [2,2]
//
//输出: [2,2]
//示例 2:
//
//输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//
//输出: [4,9]

//思路：map存储并进行遍历，数组1map值++  数组2--
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	nums := []int{}
	lenMax:= len(nums1)
	if len(nums2) > lenMax {
		lenMax =  len(nums2)
	}
	mapNum := make(map[int]int, lenMax)

	for _, v := range nums1 {
		mapNum[v]++
	}

	for _, v := range nums2 {
		if mapNum[v] > 0 {
			nums = append(nums, v)
			mapNum[v]--
		}
	}
	return nums
}

// 对于两个已经排序好数组的题，我们可以很容易想到使用双指针的解法~
func intersect2(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	nums := []int{}
	index1,index2:=0,0
	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] == nums2[index2] {
			nums = append(nums, nums1[index1])
			index1++
			index2++
		}else if nums1[index1] >  nums2[index2] {
			index2++
		}else {
			index1++
		}
	}
	return nums
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2))
	nums1 = []int{4,9,5}
	nums2 = []int{9,4,9,8,4}
	fmt.Println(intersect(nums1, nums2))
	nums1 = []int{1,2,3,4,4,13}
	nums2 = []int{1,2,3,9,10}
	fmt.Println(intersect2(nums1, nums2))
}
