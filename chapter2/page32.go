
package main

import (
	"fmt"
	"os"
	"strconv"
	"goprogramming/chapter2/tempconv"
)
func main(){
	for _,arg:=range os.Args[1:]{
		t,err:=strconv.ParseFloat(arg,64)
		if err!=nil {
			fmt.Fprintf(os.Stderr,"page32:%v\n",err)
			os.Exit(0)

		}


		f:=tempconv.Fahrenheit(t)
		c:=tempconv.Celsius(f)
		fmt.Printf("%s=%s,%s=%s\n",f,tempconv.FTOC(f),c,tempconv.CTOF(c))
	}
}

