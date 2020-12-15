package main

import (
	"fmt"
	"math"
)

func sumMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sum,max:=math.MinInt64,math.MinInt64
	for _,v :=range arr {
		if sum < 0 {
			sum = v
		}else{
			sum += v
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func main(){

	fmt.Println(sumMax([]int{1,2,3,4,-5,4,-2}))

}
