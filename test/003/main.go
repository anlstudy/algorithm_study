package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main(){
	res,err:=getMaxUserByTime("user.log",5)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}


// 在线统计用户数
func getMaxUserByTime(file string,time int) (int,error){
	change,online:=make([]int,86400),make([]int,86400)
	f,err:=os.Open(file)
	defer f.Close()
	if err!=nil{
		return 0, err
	}
	br:=bufio.NewReader(f)
	for {
		a,_,c:=br.ReadLine()
		if c == io.EOF{
			break
		}
		row:=strings.Split(string(a),",")
		up,_:=strconv.Atoi(row[0])
		down,_:=strconv.Atoi(row[1])
		change[up]++
		change[down]--
	}
	online[0] = change[0]
	for i:=1;i<=time;i++{
		online[i] = online[i-1]+change[i]
	}
	return online[time],err
}