package main

import "fmt"

//第64题：最小路径和
//给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//说明：每次只能向下或者向右移动一步。
//
//示例:
//
//输入:
//[
//[1,3,1],
//[1,5,1],
//[4,2,1]
//]
//输出: 7
//解释: 因为路径 1→3→1→1→1 的总和最小。

func minPathSum(grid [][]int) int {
	if len(grid)==0{
		return 0
	}
	dp:=make([][]int,len(grid))
	for i,arr:=range grid{
		dp[i] = make([]int,len(arr))
	}
	dp[0][0] = grid[0][0]
	for i:=1;i<len(grid);i++{
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j:=1;j<len(grid[0]);j++{
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i:=1;i<len(grid);i++{
		for j:=1;j<len(grid[i]);j++{
			dp[i][j]=min(dp[i-1][j],dp[i][j-1])+grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[len(grid)-1])-1]
}

func minPathSum2(grid [][]int) int {
	l := len(grid)
	if l < 1 {
		return 0
	}
	for i := 0; i < l; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j != 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else if i !=0 && j != 0 {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[l-1][len(grid[l-1])-1]
}


func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main(){

	fmt.Println(minPathSum([][]int{
		[]int{1,3,1},
		[]int{1,5,1},
		[]int{4,2,1},
	}))

}
