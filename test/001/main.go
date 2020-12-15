package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var stack []string
var tmp string

func main(){
	f,err:=os.Open("./test.xml")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer f.Close()
	br := bufio.NewReader(f)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		tmp = tmp+ strings.Trim(string(a)," ")
	}
	fmt.Println((format(tmp)))
}


func format(tmp string) bool {
	if len(tmp) == 0 {
		return true
	}

	for i:=0;i<len(tmp);i++{
		if string(tmp[i]) == "<" && i+1 < len(tmp) && string(tmp[i+1])!="/"{
			for j:=i+1;j<len(tmp);j++{
				if string(tmp[j]) == ">"{
					 add:=string(tmp[i+1:j])
					 stack = append(stack,add)
					 i=j
					 break
				}
			}
		}

		if string(tmp[i]) == "<" && i+1 < len(tmp) && i+2 < len(tmp) && string(tmp[i+1])=="/"{
			for j:=i+1;j<len(tmp);j++{
				if string(tmp[j]) == ">"{
					del:=string(tmp[i+2:j])
					i=j
					if del != stack[len(stack)-1]{
						return false
					}else{
						stack  = stack[:len(stack)-1]
						if len(stack) == 0 {
							if i<len(tmp)-1{
								return false
							}
						}
					}
					fmt.Println(stack)
					break
				}
			}
		}
	}
	return true
}
