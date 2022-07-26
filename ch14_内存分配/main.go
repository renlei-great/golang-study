package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func main() {
	/*
	go语言的内存由go语言本身去管理，go语言管理的虚拟内存空间分两部分：堆和栈内存，栈内存主要由go语言来管理，开发者无法干涉太多，堆内存才是开发者的舞台
	一个程序大部分的内存分配到了堆内存中，常说的go内存垃圾回收针对的是堆内存的垃圾回收

	变量在声明后，会被分配一个内存空间，然后值是零值，也可以在声明时初始化,初始化后其实就已经赋值了
		var s string = "rl"
		s := "rl"

	new 函数，可以通过new函数申请一块相应类型的内存空间，然后返回这个内存空间的地址，也就是返回一个指针类型
		作用就是根据传入的类型申请一块内存，然后返回指向这块内存的指针，指针指向的就是这个类型的零值

	make和new的作用相同，但它是去申请一些复杂类型，slice，chan, map这三种内置类型的创建和初始化
	 */
	var s string
	var sp *string
	sp = new(string)
	fmt.Println(sp)
	fmt.Printf("%p\n", &	s)

	pp := NewPerson("re", 19)
	fmt.Println(pp.name, pp.age)
}

func NewPerson(name string, age int) *person {
	p := new(person)
	p.name = name
	p.age = age
	return p
}
