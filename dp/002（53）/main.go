package main



//第53题：最大子序和
//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。


func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max,subMax:=nums[0],nums[0]
	for i:=1;i<len(nums);i++{

		subMax = subMax + nums[i]
		if subMax < nums[i]{
			subMax = nums[i]
		}

		if subMax > max {
			max = subMax
		}
	}
	return max
}

func main(){

}
