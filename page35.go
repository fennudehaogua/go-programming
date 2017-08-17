package main

import "fmt"

func main(){
	x:="hello!"
	for i:=0;i<len(x);i++{
		x:=x[i]
		if x<'z'&&x>'a'{
			x=x-'a'+'A'

		}
		fmt.Printf("%c",x)



	}
}