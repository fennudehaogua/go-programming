<<<<<<< HEAD
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

)

func main(){
	counts:=make(map[string]int)
	for _,filename := range os.Args[1:]{
		data,err:= ioutil.ReadFile(filename)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"dup3:%v\n",err)
			continue
		}
		for _,line:=range strings.Split(string(data),"\n"){
			counts[line]++
		}
	}
	for line,n:=range counts{
		if n>1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
=======
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

)

func main(){
	counts:=make(map[string]int)
	for _,filename := range os.Args[1:]{
		data,err:= ioutil.ReadFile(filename)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"dup3:%v\n",err)
			continue
		}
		for _,line:=range strings.Split(string(data),"\n"){
			counts[line]++
		}
	}
	for line,n:=range counts{
		if n>1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
>>>>>>> 24181879451b8db8a6db4d439894d3ade71ca404
}