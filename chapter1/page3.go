package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	var s,sep string
	for i := 1; i < len(os.Args); i++ {
		s+= sep+os.Args[i]
		sep = " "
	}
	fmt.Println(s)


	fmt.Println("===============")



	for j := 1; j<len(os.Args);j++  {
		fmt.Printf("Args[%d]=%s\n",j,os.Args[j])
	}
	fmt.Println("===============")

	for key,value:=range os.Args{
		fmt.Printf("Args[%d]=%s\n",key,value)
	}
	fmt.Println("===============")


	fmt.Println(strings.Join(os.Args[0:]," "))
}

