package main

import "fmt"

//请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
//
//示例1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

func lengthOfLongestSubstring(s string) int {
	strMap:=make(map[string]int,len(s))
	max,tmp:=0,0
	for i:=0;i<len(s);i++{
		if value,ok:=strMap[string(s[i])];ok{
			fmt.Println("碰撞",tmp,i,string(s[i]))
			if i - value <= tmp {
				tmp = i-value
			}else{
				tmp++
			}
		}else{
			tmp++
			fmt.Println("无碰撞",tmp,i,string(s[i]))
		}
		fmt.Println(tmp)
		fmt.Println(max)
		if tmp > max {
			max = tmp
		}

		strMap[string(s[i])] = i
	}
	return max
}


func main(){
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

