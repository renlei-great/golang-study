package main

import "fmt"

func main() {
	/*
	本节主要认识参数传递的时候，值和引用及指针之间的区别
	*** 在go语言中，函数的参数传递只有值传递，而且传递的实参数都是原始数据的一份拷贝

		修改参数需要传递的是指针类型，因为指针的值是存着原始数据的内存地址，这样就算被拷贝了一份也同样可以找到原始数据的位置并进行修改
		值类型在传递的时候是重新拷贝了一份放在了新的内存中，这样就算修改各种操作都和原始内存是没有关系的
			浮点型，整型，字符串，布尔，数组，都是值类型

		可称为引用类型的
			指针类型的变量的值保存了数据对应的内存地址，所以在函数传递中拷贝一份变量的值，还是拷贝了内存地址，所以就可以修改值了
			像map，chan当用字面量或者make函数方法创建的时候都会由go语言自动实现去调用runtime.makemap或runtime.makechan,这个函数返回的是一个指针类型，所以传递可称为引用类型的变量是可以修改原始值的
				但是严格来说，go语言是没有引用类型的
				map, chan, 函数， 接口， slice 都称为引用类型
	 */

	p := person{name: "rl", age: 18}
	modifyPerson(&p)
	fmt.Println(p.name, p.age)

	m := make(map[string]int)
	m["rl"] = 18
	fmt.Println("rl:", m["rl"])
	fmt.Printf("m的内存地址：%p\n", m)
	modifyMap(m)
	fmt.Println("rl:", m["rl"])

	s := make([]int, 5, 5)
	fmt.Println(s)
	modifySlice(s)
	fmt.Println(s)

}

func modifySlice(s []int)  {
	s[0] = 1

}

func modifyMap(p map[string]int)  {
	p["rl"] = 20
	fmt.Printf("p的内存地址：%p\n", p)

}

func modifyPerson(p *person)  {
	p.name = "站"
	p.age = 1
}

type person struct {
	name string
	age int
}
