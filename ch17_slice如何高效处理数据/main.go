package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

type slice struct {
	D uintptr
	L int
	C int
	E int
}

func main() {
	/*
	这节主要学习根据熟悉底层设计后，根据底层设计配合指针提高效率

	数组：
		由两部份组成，一部分是数组的大小，一部分是数组的元素类型
		只有数组的长度和元素类型都相同这个时候两个数组才算同一类型数组
		数组必须要提前定好他的长度，函数间的参数传递是值传递，传来传去这样带来的问题就是会一份份的拷贝这个数组
		但是数组是非常重要的底层数据结构，比如slice切片地城数据存储在数组中

	slice：
		是对数组的抽象和封装
		底层是数组存储数据，但是可以动态的添加元素，容量不足会自动扩容，可以理解为是动态数组
		append 自动扩容原理是新创建一个数组，然后把原来的元素拷贝到新数组，然后返回一个指向新数组的切片
		数据结构：
			type SLiceHeader struct {
				Data uintptr
				Len int
				Cap int
			}
		多个切片可以共用一个底层数组，达到节约内存

	数组和切片的取值赋值操作都高效，因为是连续的内存操作

	string和[]byte 类型转换
		使用string() 和 []byte() 进行强制类型转换,大体过程
			1。 go语言先分配值的内存
			2。 进行复制内容
			3。 然后赋值给新的结构体
			4。 类型转换完成
		优化：
			通过unsafe 直接修改结构体，不去重新创建新的数组，共用一个data，达到优化内存
			但是因为string是只读内存，所以使用此优化方式[]byte去执行修改操作时会报错
	 */


	s := "renlei niupbi "
	fmt.Printf("s的内存地址是：%d", (*reflect.StringHeader)(unsafe.Pointer(&s)), s, "\n")
	//b := []byte(s)
	//fmt.Printf("b的内存地址是：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)), b)
	rs := *(*reflect.SliceHeader)(unsafe.Pointer(&s))
	rs.Cap = rs.Len
	b1 := *(*[]byte)(unsafe.Pointer(&rs))
	fmt.Printf("b1的内存地址是：%d", (*reflect.SliceHeader)(unsafe.Pointer(&b1)).Data, b1, "\n")
	//s3 := string(b)
	//fmt.Printf("s3的内存地址是：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)))
	s4 := *(*string)(unsafe.Pointer(&b1))
	fmt.Printf("s4的内存地址是：%d", (*reflect.StringHeader)(unsafe.Pointer(&s4)), s4, "\n")

	ch := make(chan string)
	go func() {
		//time.Sleep(time.Second)
		fmt.Println("111")
		ch <- "sss"
		fmt.Println("222")
	}()

	fmt.Println(ch)
	chp := *(*reflect.ChanDir)(unsafe.Pointer(&ch))
	time.Sleep(time.Second)
	fmt.Println("333")
	a := <-ch
	fmt.Println("444")
	fmt.Println(chp, a)
	time.Sleep(time.Second)


}

func test2() {
	// 验证数组在函数传递时会拷贝一份值，而切片不会重新拷贝数组

	a1:= [2]string{"rl", "ll"}
	fmt.Printf("函数main数组指针：%p\n", &a1)
	arrayF(a1)
	s1:=a1[0:1]
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	sliceF(s1)

}

func arrayF(a [2]string) {
	fmt.Printf("函数array数组指针：%p\n", &a)
}

func sliceF(s []string) {
	fmt.Printf("函数sliceF Data: %d\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)

}

func test1(){
	// 获取切片的结构体
	ss := []string{"el", "gq"}
	ss = append(ss, "ff", "aa")
	fmt.Println(ss)
	fmt.Println(len(ss), cap(ss))

	a1 := [2]string{"rl", "rll"}
	s1 := a1[0:1]
	s2 := a1[:]
	s1 = append(s1, "ll")
	s1u := unsafe.Pointer(&s1)
	s1p := (*slice)(unsafe.Pointer(&s1))
	fmt.Println(s1p, s1u)
	fmt.Println((*slice)(unsafe.Pointer(&s1)).D)
	fmt.Println((*slice)(unsafe.Pointer(&s2)).D)
}
