package main

import (
	"fmt"
	"unsafe"
)

type person struct {
	Name string
	Age int
}

func main() {
	/*
	指针类型转换
		普通指针类型 *T
		unsafe.Pointer  一种特殊意义的指针，可以表示任意类型地址
		uintptr  指针类型, 也可以表示任意地址并且此指针类型可以进行加减去移动指向的内存，这样就可以进行偏移计算进行加减去找到对应的内存进行修改

	转换规则：普通可以转unsafe.Pointer，unsafe.Pointer可以转普通，unsafe.Pointer可以转uintptr， uintptr可以转unsafe.Pointer，uintptr不能转普通，普通不能转uintptr
		1。 任何类型的*T 都可以转unsafe.Pointer
		2。 unsafe.Pointer 也可以转任何类型的 *T
		3.  unsafe.Pointer 可以转 uintptr
		4.  uintptr 也可以转换为 unsafe.Pointer

	unsafe.Sizeof 可以返回一个类型所占的内存大小
	 */

	p := new(person)
	// name是person的第一个字段，即可直接通过指针修改
	pName := (*string)(unsafe.Pointer(p))
	*pName = "rll"
	fmt.Println(p)
	// age并不是第一个，所以需要进行偏移，才能找到age的内存，进行修改
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age)))
	a :=unsafe.Pointer(p)
	b := uintptr(unsafe.Pointer(p))
	c := unsafe.Offsetof(p.Age)
	fmt.Println(p, a, b, c)
	*pAge = 20
	fmt.Println(p)
	fmt.Println(unsafe.Sizeof("rrl"))
	fmt.Println(unsafe.Alignof(p))



	// 这个示例通过直接找到内存，然后在修改内存中的值，而不是通过变量名
	/*
	1。 先使用 new 函数声明一个 *person 类型的指针变量 p。
	2。 然后把 *person 类型的指针变量 p 通过 unsafe.Pointer，转换为 *string 类型的指针变量 pName。
	3。 因为 person 这个结构体的第一个字段就是 string 类型的 Name，所以 pName 这个指针就指向 Name 字段（偏移为 0），对 pName 进行修改其实就是修改字段 Name 的值。
	4。 因为 Age 字段不是 person 的第一个字段，要修改它必须要进行指针偏移运算。所以需要先把指针变量 p 通过 unsafe.Pointer 转换为 uintptr，这样才能进行地址运算。既然要进行指针偏移，那么要偏移多少呢？这个偏移量可以通过函数 unsafe.Offsetof 计算出来，该函数返回的是一个 uintptr 类型的偏移量，有了这个偏移量就可以通过 + 号运算符获得正确的 Age 字段的内存地址了，也就是通过 unsafe.Pointer 转换后的 *int 类型的指针变量 pAge。
	5。 然后需要注意的是，如果要进行指针运算，要先通过 unsafe.Pointer 转换为 uintptr 类型的指针。指针运算完毕后，还要通过 unsafe.Pointer 转换为真实的指针类型（比如示例中的 *int 类型），这样可以对这块内存进行赋值或取值操作。
	6。 有了指向字段 Age 的指针变量 pAge，就可以对其进行赋值操作，修改字段 Age 的值了。
	 */

}


func test1() {
	// 简单测试万能指针的使用
	i:= 10
	ip := unsafe.Pointer(&i)
	var fp *float64 = (*float64)(ip)
	*fp = *fp * 3
	fmt.Println(i)
}

