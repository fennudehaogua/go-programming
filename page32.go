<<<<<<< HEAD
package main

import (
	"fmt"
	"os"
	"strconv"
	"goprogramming/tempconv"
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
=======
package main

import (
	"fmt"
	"os"
	"strconv"
	"goprogramming/tempconv"
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
>>>>>>> 24181879451b8db8a6db4d439894d3ade71ca404
