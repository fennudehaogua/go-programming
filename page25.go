package main

import "fmt"

func main(){
	p:=new(int)
	fmt.Println(*p)
	fmt.Println(newInt1()==newInt1())
	fmt.Println(newInt1()==newInt2())

	fmt.Println(*newInt1())
	fmt.Println(*newInt1())
	fmt.Println(*newInt1())

	fmt.Println(*newInt2())
	fmt.Println(*newInt2())
	fmt.Println(*newInt2())



	var x,y=1,2
	x,y=y,x
	fmt.Println(x,y)

	var a [2]int=[2]int{1,2}
	a[0],a[1]=a[1],a[0]
	fmt.Println(a[0],a[1])


	fmt.Println(maxGYS(12,18))
	fmt.Println(fib(10))
}


func newInt1() *int{
	return new(int)
}

func newInt2() *int{
	var num int
	return &num
}

func maxGYS(a int,b int) int{
	for b!=0{
		a,b=b,a%b
	}
	return a
}


func fib(n int) int{
	x,y:=0,1
	for i:=0;i<n;i++{
		x,y=y,x+y
	}
	return x
}