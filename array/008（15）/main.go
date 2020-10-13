package main

import (
	"fmt"
	"sort"
)

//该题为 二数之和 的进阶版本，当然还有一个进阶版本为 四数之和。我们将会一一进行分析！
//
//第15题：三数之和
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
//示例：
//
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//满足要求的三元组集合为：
//[
//[-1, 0, 1],
//[-1, -1, 2]
//]

func threeSum(nums []int) [][]int {
	fmt.Println(nums)
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for index:=0;index<n;index++{
		if index > 0 && nums[index] == nums[index-1]{
			continue
		}
		end:=n-1
		target:= -1 * nums[index]
		for start:=index+1;start<n;start++{
			if start > index+1 && nums[start] == nums[start-1]{
				continue
			}

			for start<end && nums[start]+nums[end] > target {
				end --
			}

			if start == end {
				break
			}

			if nums[start] + nums[end] == target {
				ans = append(ans,[]int{nums[index],nums[start],nums[end]})
			}
		}
	}
	return ans
}


func main(){
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
