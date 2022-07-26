package main

import (
	"fmt"
	"strings"
)

func main() {

	var i int = 10
	fmt.Println(i)

	var a = 10
	fmt.Println(a)

	var (
		b = 10
		c = 10
	)
	fmt.Println(b, c)

	var f32 float32 = 2.2

	var f64 float64 = 10.3456

	fmt.Println("f32 is", f32, ",f64 is", f64)

	var s1 string = "Hello"

	var s2 string = "世界"

	fmt.Println("s1 is", s1, ",s2 is", s2)

	fmt.Println("s1+s2=", s1+s2)

	s1 += s2

	fmt.Println("s1:", s1)

	var zi int

	var zf float64

	var zb bool

	var zs string

	fmt.Println(zi, zf, zb, zs)

	ii := 10

	// bf:=false

	// s1:="Hello"

	pi := &ii

	fmt.Println(*pi)

	ss1 := "abcdef"
	ss2 := "defg"

	fmt.Println(strings.HasSuffix(ss1, ss2))
	/*

		整型:
			有符号整型：类似int8, 特殊的是int 会随着硬件进行变化
			无符号整型：类似uint8, 特殊的是uint 会随着硬件进行变化
			go 语言中还有一种字节类型byte，它其实是等价于uint8类型，可以理解为uint8的别名

		浮点数：
			go语言中提供了两种精度的浮点数，分别是float32,float64,一般常用的是float64,精度高，误差会小一些

		布尔值：真或假

		字符串string :

		零值：
			零值就是变量的默认值，在go语言中，如果我们声明了一个变量，但是没有对其进行初始化
			，那么go语言就会自动出实话其值为对应类型的零值，比如数字是0，bool是false，字符串是空
	*/

	/* 变量

	变量简短声明：
		可以使用:=进行声明变量
			变量名 := 表达式

	指针：
		在go语言中，指针对应的就是变量在内存中的存储位置，指针的值其实就是变量的内存地址
		通过 & 可以获取一个变量的地址，也就是指针

	赋值：值给变量赋值

	*/

	/*

		常量：
			go语言中只允许布尔型、字符串、数字类型这些基础类型作为常量
			const name =  "飞雪无情"

			iota 是一个常量生成器
				const(
					one = 1
					two = 2
					three =3
					four =4
				)
				可以使用下面代码代替
				const(
					one = iota+1
					two
					three
					four
				)
	*/

	/*

		字符串

				字符串和数字互转使用strconv
					也可以使用 float64(), int()

				strings 标准包
					//判断s1的前缀是否是H

					fmt.Println(strings.HasPrefix(s1,"H"))

					//在s1中查找字符串o

					fmt.Println(strings.Index(s1,"o"))

					//把s1全部转为大写

					fmt.Println(strings.ToUpper(s1))


	*/

	/*

	 */
	/*

	 */
	/*

	 */
}
