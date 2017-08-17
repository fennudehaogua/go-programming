
package main

import (
	"fmt"
	"os"
	"strconv"
	"goprogramming/chapter2/lengthconv"
)

func main(){
	for _,arg:=range os.Args[1:]{
		t,err:=strconv.ParseFloat(arg,64)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"page32test:%v\n",err)
			os.Exit(1)
		}

		i:=lengthconv.Incun(t)
		m:=lengthconv.Meter(t)
		fmt.Printf("%s=%s,%s=%s\n",i,lengthconv.ITOM(i),m,lengthconv.MTOI(m))
	}
}

