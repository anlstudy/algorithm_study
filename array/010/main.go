package main

import (
	"fmt"
	"strconv"
)

func main(){
	arr:=[]int{99,567,9,88}
	sort(arr)
	fmt.Println(arr)

}

func sort(arr []int ){
	for i:=0;i<len(arr);i++{
		for j:=i;j<len(arr);j++{
			if max(arr[i],arr[j]){
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}

}

func max(a int ,b int) bool{
	first:= strconv.Itoa(a)+strconv.Itoa(b)
	second:= strconv.Itoa(b)+strconv.Itoa(a)
	fir,_:= strconv.Atoi(first)
	sec,_:= strconv.Atoi(second)
	if fir > sec {
		return true
	}
	return false
}
