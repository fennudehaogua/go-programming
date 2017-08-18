package main

import (
	"fmt"
	"math"
)

const (
	width,height=600,320
	cells = 100
	xyrange = 30.0
	xyscale = width/2/xyrange
	zscale = height*0.4
	angle = math.Pi/6

)
var sin30,cos30 = math.Sin(angle),math.Cos(angle)
func main(){

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg'"+
		"style='stroke:grey;fill:white;stroke-width:0.7'"+
		"width='%d' height='%d'>",width,height)



	for i:=0;i<cells;i++{
		for j:=0;j<cells;j++{
			ax,ay:=corner(i+1,j)
			bx,by:=corner(i,j)
			cx,cy:=corner(i,j+1)
			dx,dy:=corner(i+1,j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",ax,ay,bx,by,cx,cy,dx,dy)

		}
	}
	fmt.Println("</svg>")
}





func corner(i int, j int) (float64, float64) {
	x:=xyrange*(float64(i)/cells-0.5)
	y:=xyrange*(float64(j)/cells-0.5)

	z := f(x,y)

	sx:=width/2+(x-y)*cos30*xyscale
	sy:=height/2+(x-y)*sin30*xyscale-z*zscale
	return sx,sy
}


func f(x float64,y float64) float64{
	r:=math.Hypot(x,y)
	return math.Sin(r)
}

//运行时碰到输出的结果cmd显示不完全。只好将输出打入一个文件中，go run page43.go>>1.txt,然后将其在网页中运行，但是结果与书中不一样，具体问题也不知道为什么