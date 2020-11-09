package main

import "fmt"

//第70题：爬楼梯
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？ **注意：**给定 n 是一个正整数。
//示例 1：
//
//输入： 2   输出： 2  解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//示例 2：
//
//输入： 3   输出： 3  解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶

//func climbStairs(n int) int {
//	if n <=2{
//		return n
//	}
//	return climbStairs(n-1)+climbStairs(n-2)
//}

func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	res:=make([]int,n)
	res[0] = 1
	res[1] = 2
	for index:=2;index<n;index++{
		res[index] = res[index-1]+res[index-2]
	}
	return res[n-1]
}


func climbStairs(n int) int {
	if n < 3 {
		return n
	}
    a,b,c := 2,1,0
	for index:=3;index<=n;index++{
		c = a+b
		b = a
		a = c

	}
	return c
}


func main(){
  fmt.Println(climbStairs(10))
}
