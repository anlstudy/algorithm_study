package main

import (
	"fmt"
	"math"
)

//
//假如要计算11元需要的面值数最小，那么，11元减去一张1元或者3元或者5元，即10元，8元，6元分别需要的面值数，取出最小，加上1，即可以得到11元最小的。
//所以，n元最小问题，是由n-1元，n-3元，n-5元三种状态里面取到的最优解。
//其中计算的递归树为：
//d(11) = min{d(10),d(8),d(6)} + 1
//递推式
//d(i) = min{d(i-v)} + 1,其中v为面值1，3，5；i为求值

func select_coin(coin[]int, total int)int{
	length:=total+1
	minCoins:=make([]int,length)
	minCoins[0] = 0
	for i := 1; i <= total ; i++ {
		minCoins[i] = math.MaxInt64
		for j:=0; j<len(coin);j++ {
			if coin[j] <= i && minCoins[i-coin[j]]+1 < minCoins[i] {
				minCoins[i] = minCoins[i-coin[j]] + 1
			}
		}
	}
	fmt.Println(minCoins)
	return minCoins[total]
}

func main(){
	fmt.Println(select_coin([]int{1,3,5},11))
}