<<<<<<< HEAD
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
=======
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
>>>>>>> 24181879451b8db8a6db4d439894d3ade71ca404
}