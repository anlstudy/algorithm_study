package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 按照2:2:3的概率返回ABC三个中 两个

func main(){
	fmt.Println(getAd())
}

func getAd()(string,string){
	rand.Seed(time.Now().Unix())
	var one,two string
	num:=rand.Intn(7)
	if num < 2 {
		one = "A"
	}else if num < 4{
		one = "B"
	}else{
		one = "C"
	}

	for{
		num=rand.Intn(7)
		if num < 2 {
			two = "A"
		}else if num < 4{
			two = "B"
		}else{
			two = "C"
		}
		if one!=two{
			break
		}
	}
	return one,two
}
