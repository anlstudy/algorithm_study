package main

import "fmt"

//题目27：移除元素
//给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
//不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
//
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//
//示例 1:
//
//给定 nums = [3,2,2,3], val = 3,
//函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
//你不需要考虑数组中超出新长度后面的元素。

func removeElement(nums []int, val int) int {
	low:=0
	for _,v:= range nums {
		if v != val {
			nums[low] = v
			low++
		}
	}
	return low
}


func removeElement2(nums []int, val int) int {
	i,j:=0,len(nums)
	for i < j {
		if nums[i] == val {
			nums[i] = nums[j-1]
			j--
		}else{
			i++
		}
	}
	return j
}

func main(){
    arr:=[]int{3,3,2,2}
	fmt.Println(removeElement(arr,3))
	fmt.Println(arr)

	arr1:=[]int{3,3,2,2}
	fmt.Println(removeElement2(arr1,3))
	fmt.Println(arr1)
}
