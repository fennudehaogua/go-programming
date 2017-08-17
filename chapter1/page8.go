
//page8.go
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	counts:=make(map[string]int)
	files:=os.Args[1:]//os.Args[0]是命令本身的名字
	if len(files)==0{//不带文件名
		countLines(os.Stdin,counts)
	}else{
		for _,arg :=range files{
			f,err:=os.Open(arg)
			if err!=nil{
				fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
				continue
			}
			countLines(f,counts)
			f.Close()
		}
	}
	for line,n :=range counts{
		if n>1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}

}



func countLines(f *os.File, counts map[string]int) {
	input:=bufio.NewScanner(f)
	for input.Scan() {
		if input.Text()=="end" {
			break
		}
		counts[input.Text()]++

	}




}

//在当前路径下新建两个文本文件输入内容，执行go run page8.go xxx.txt xxx.txt

//如果直接执行go run page8.go则功能是检查输入的内容