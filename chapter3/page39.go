package main
import (
	"fmt"
	"math"
)

const (
	x uint8 = 1<<1 | 1<<5
	y uint8 = 1<<1 | 1<<2


)
//定义为const会报错
var f1 float32 = 1<<24
var u uint8 =255
var i int8 = 127
func main(){
	fmt.Println(u,u+1,u*u)
	fmt.Println(i,i+1,i*i)

	fmt.Printf("%08b\n",x)
	fmt.Printf("%08b\n",y)


	fmt.Printf("%08b\n",x&y)
	fmt.Printf("%08b\n",x|y)
	fmt.Printf("%08b\n",x^y)//异或，不同的取1
	fmt.Printf("%08b\n",x&^y)//如果y的某位是1，则结果的对应位为0，否则与x的相同

	for i:=uint(0);i<8;i++{
		if x&(1<<i)!=0{
			fmt.Println(i)
		}
	}



	fmt.Printf("%08b\n",x<<1)
	fmt.Printf("%08b\n",x>>1)



	fmt.Println(f1==f1+1)


	for x:=0;x<8;x++{
		fmt.Printf("x=%d e的%d次方=%8.3f\n",x,x,math.Exp(float64(x)))
	}
}
