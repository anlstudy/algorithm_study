package main

import "fmt"

//第66题：加一
//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。你可以假设除了整数 0 之外，这个整数不会以零开头。
//
//示例 1:
//
//输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。
//示例 2:
//
//输入: [4,3,2,1]
//输出: [4,3,2,2]
//解释: 输入数组表示数字 4321。
//
//题目分析：
//
//根据题目，我们需要加一！没错，加一很重要。因为它只是加一，所以我们会考虑到两种情况：
//
//普通情况，除9之外的数字加1。
//
//特殊情况，9加1。（因为9加1需要进位）
//
//所以我们只需要模拟这两种运算，就可以顺利进行求解

func plusOne(digits []int) []int {
	length:=len(digits)
	if length == 0 {
		return  digits
	}

	if digits[length-1]  != 9 {
		digits[length-1] = digits[length-1]+1
		return digits
	}

	ex:= 1
	for ex == 1 && length > 0 {
		if digits[length-1]+1 == 10 {
			digits[length-1] = 0
			length --
		}else {
			ex = 0
		}
	}

	if ex == 0 {
		digits[length-1] = digits[length-1] + 1
		return digits
	}else {
		return append([]int{1},digits...)
	}
}

func main(){
	fmt.Println(plusOne([]int{9,9,9}))
}
