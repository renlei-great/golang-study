package main

import (
	"errors"
	"fmt"
)

type Age uint

func (age Age) String() {
	fmt.Println("the age is", age)
}

func (a Age) Print(x int) {
	fmt.Println("the age is, x", a, x)
}

func (age *Age) Modify() {
	*age = Age(30)
}

func sum(a, b int) (int, error) {

	if a < 0 || b < 0 {

		return 0, errors.New("a或者b不能是负数")

	}

	return a + b, nil

}

func sum1(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}
	return sum
}

func main() {

	// v, e := sum(-1, 2)
	v := sum1(-1, 2, 3, 4, 5)
	fmt.Println(v)

	// 包级函数，不管是自定义的函数还是使用到的函数Println 都是属于某一个包的，也就是package ,sum 属于main包，Println 数据fmt包
	// 同一个包的函数怎么样都可以调用，就算是私有的(函数名首字母小写)
	// 如果是不同包的函数要被调用，函数的作用域必须是公有的，也就是函数首字母大写，比如Println
	// 首字母小写是私有函数，只可以在一个包中调用
	// 首字母大写是公有函数，不同的包都可以调用
	// 任何一个函数都会从属于一个包

	// 匿名函数和闭包， 通过匿名函数实现闭包
	sum2 := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum2(1, 2))

	cl := colsure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())

	a := Age(25)
	a.String()
	a.Print(1)
	a.Modify()
	a.Print(2)
	// 可以使用值传递，也可以使用指针类型传递，都是可以的，到了接口那边，go会根据定义自动做转化
	(&a).Modify()
	fmt.Println(a + 1)

	// 方法可以被赋值，但是赋值以后调用的时候他的第一个参数是接收者
	aStr := Age.Print
	aStr(10, 1)
}

func colsure() func() int {
	// go中，函数也是一种类型，可以被用来声明函数类型变量，作为返回值进行返回

	i := 0
	return func() int {
		i++
		return i
	}
}
