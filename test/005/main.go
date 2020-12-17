package main

import "fmt"

// 从做左到右升序，从上到下降序，找元素i
func main(){
	arr := [][]int{
		{9,11,17},
		{3,7,13},
		{1,5,11},
	}
	fmt.Println(isExist(arr,3,3,5))
}

func isExist(arr [][]int,m,n,num int) bool {
	i,j:=0,0
	count:=0
	for i<m&&j<n{
		if i<m && j<n && arr[i][j] == num {
			fmt.Println(count)
			return true
		}
		if i<m && j<n && arr[i][j] < num {
			j++
			count++
		}
		if i<m && j<n && arr[i][j] > num {
			i++
			count++
		}
	}
	fmt.Println(count)
	return false
}
