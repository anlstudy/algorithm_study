package main


//第387题：字符串中的第一个唯一字符
//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1 。
//案例:
//
//s = "leetcode"
//返回 0.
//
//s = "loveleetcode",
//返回 2.

func firstUniqChar(s string) int {
	var arr [26]int
	for i,k := range s {
		arr[k - 'a'] = i
	}
	for i,k := range s {
		if i == arr[k - 'a']{
			return i
		}else{
			arr[k - 'a'] = -1
		}
	}
	return -1
}

func main(){

}
