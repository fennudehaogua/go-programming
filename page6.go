<<<<<<< HEAD
package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){
	counts :=make(map[string]int)
	input:= bufio.NewScanner(os.Stdin)
	for input.Scan(){
		//控制循环退出
		if input.Text() == "end" { break }
		counts[input.Text()]++
	}
	for line,n :=range counts{
		if n>1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}

=======
package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){
	counts :=make(map[string]int)
	input:= bufio.NewScanner(os.Stdin)
	for input.Scan(){
		//控制循环退出
		if input.Text() == "end" { break }
		counts[input.Text()]++
	}
	for line,n :=range counts{
		if n>1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}

>>>>>>> 24181879451b8db8a6db4d439894d3ade71ca404
