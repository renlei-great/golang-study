package main

import "fmt"

func main() {
	/*
	指针的值是一个内存地址，内存地址通常是一串16进制数
		普通变量的存储的是值
		指针存储的是一个内存地址

	指针接受者
		1。 如果接受者类型是map,slice,channel 这类引用类型，不使用指针
		2。 如果需要修改接受者，那么需要使用指针
		3。 如果接受者是比较大的类型，可以考虑使用指针，因为内存拷贝廉价，效率高

	什么情况下使用指针
		指针的两大好处
			1。 可以修改指向数据的值
			2。 在变量赋值，参数传值的时候可以节省内存

		1。 不对map, slice, channel 引用类型使用指针
		2。 如果需要修改方法接受者内部的值的时候使用指针
		3。 如果需要修改参数的值或者内部数据时，使用指针
		4。 如果比较大的结构体，拷贝内存占用多的情况使用指针
		5。 像int. bool 这样比较小的数据没有必要使用指针
		6。 如果需要并发安全，尽可能不使用指针，使用指针一定要保证并发安全
		7。 指针别嵌套使用

	 */
	name := "rl"
	nameP := &name
	fmt.Println("name:", name)
	fmt.Println("nameP:", nameP)

	nameV := *nameP
	fmt.Println("nameP指向的值",nameV)

	*nameP = "rl1"

	fmt.Println("name:", name)
	fmt.Println("nameP:", nameP)

	age := 18
	modifyAge(&age)
	fmt.Println(age)

	// 此段可以确定值类型实现接口，指针类型也实现了接口
	add := address{province: "aaa", city: "bbb"}
	printString(add)
	printString(&add)

	var si fmt.Stringer = address{province: "aaa", city: "bbb"}
	printString(si)
	sip := &si
	printString(*sip)


}

func printString(s fmt.Stringer)  {
	fmt.Println(s.String())
}

type address struct {
	province string
	city string
}

func (addr address) String() string {
	return fmt.Sprintf("我是testFunc %s, %s", addr.province, addr.city)
}

func modifyAge(age *int){
	*age = 20

}
