package main

import "fmt"

func main()  {
	f:=212.0
	c:=(f-32)*5/9
	fmt.Printf("version1:%g摄氏度=%g华氏度\n",c,f)
	fmt.Printf("version2:%f摄氏度=%f华氏度\n",changeFromFtoC(f),f)
}

func changeFromFtoC(f float64)  float64{
	return (f-32)*5/9;
}
//%g用于打印浮点型数据时，会去掉多余的零，至多保留六位有效数字（不同于%e的默认保留小数点后6位）
//当%g用于打印超过6位的浮点型数据时，因为精度问题，%f不得不输出一个不精确的超过六位的数字，%e也是同样，而%g此时会选择%e格式进行输出，并且按第一条要求，去掉多余的零，并且四舍五入到6位数字。这是《C Primer Plus》中所说的超过精度的时候的情况。  （可见，这个6位，是按float类型精度来计算的。）
//当一个数字的绝对值很小的时候，要表示这个数字所需要的字符数目就会多到让人难以接受。举例而言，如果我们把π*10^-10写作0.00000000000314159就会显得非常丑陋不雅，反之，如果我们写作3.14159e-10，就不但简洁而且易读好懂。当指数是-4时，这两种表现形式大小相同。对于比较小的数值，除非该数的指数小于或者等于-5，%g才会采用科学技术发来表示，即，以%e的格式进行输出。