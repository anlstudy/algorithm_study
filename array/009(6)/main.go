package main

import (
	"fmt"
)

//第6题：Z 字形变换
//将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
//L   C   I   R
//E T O E S I I G
//E   D   H   N
//
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
//
//
//请你实现这个将字符串进行指定行数变换的函数：
//
//string convert(string s, int numRows);
//示例 1:
//
//输入: s = "LEETCODEISHIRING", numRows = 3
//输出: "LCIRETOESIIGEDHN"
//示例 2:
//
//输入: s = "LEETCODEISHIRING", numRows = 4
//输出: "LDREOEIIECIHNTSG"
//
//解释:
//L     D     R
//E   O E   I I
//E C   I H   N
//T     S     G

func convert(s string, numRows int) string {
	if numRows == 0 || numRows == 1 {
		return s
	}
	res:=make([][]string,numRows)
	for i:=0;i<numRows;i++{
		res[i] = make([]string,len(s))
	}
	index,x,y:=0,0,0
	flag:=true
	for index < len(s){
		res[x][y] = string(s[index])
		index++
		if flag {
			if x == numRows-1 {
				x--
				y++
				flag = !flag
				continue
			}else{
				x++
			}
		}

		if !flag{
			if x == 0 {
				x++
				flag=!flag
				continue
			}else{
				x--
				y++
			}
		}
	}
	resStr:=""
	for i:=0;i<numRows;i++{
		for j:=0;j<len(s);j++{
			if res[i][j]!=""{
				resStr = resStr+res[i][j]
			}
		}
	}
	return resStr
}

func main(){
	fmt.Println(convert("LEETCODEISHIRING",3))
}
