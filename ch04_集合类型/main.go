package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	array := [5]string{"a", "b", "c", "d", "e"}
	// array := [...]string{"a", "b", "c", "d", "e"}  // 可以省略数组的长度，这个时候go会自动根据大括号中的元素推导长度

	fmt.Println(array[2])

	// 意思是初始化索引1的值为b，下标3的值为d，这个也可以省略【5】为 [...]，但是长度会变成4，
	// 其他没有初始化到的下标为“”
	array1 := [5]string{1: "b", 3: "d"}
	fmt.Println(array1[3])

	// 下面是go的新型循环， for range
	for i, v := range array {

		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)

	}

	// 切片

	slice := array[2:5]
	fmt.Println("slice:", slice)
	slice[1] = "f"

	fmt.Println(array)

	// 切片声明, 4是切片的长度，8是cap切片的容量，使用append 会用到空闲的容量5，如果满了会进行扩容
	// slice1 := make([]string, 4, 8)
	// 切片的另一种声明方式
	slice1 := []string{"a", "b", "c"}
	fmt.Println(len(slice1), cap(slice1))
	// 可以使用内置函数append 进行元素的追加，会返回一个新的切片
	slice2 := append(slice1, "d")
	fmt.Println(slice2)
	slice2 = append(slice1, "e", "f")
	fmt.Println(slice2)
	slice2 = append(slice1, slice...)
	fmt.Println(slice2)
	// ****小技巧，如果生成的时候长度和容量相同，当追加的时候就会生成新的低层数组

	fmt.Println("遍历切片")
	for i, v := range slice2 {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}

	// 创建map
	// nameAgeMap := make(map[string]int)
	// nameAgeMap["rl"] = 20
	nameAgeMap := map[string]int{"rl": 20} // 使用此方法，大括号中的内容可以省略，但是大括号不可以省略
	fmt.Println(nameAgeMap)
	// go语言中的map 可以获取不存在的k-v，如果不存在会返回该类型的零值，比如int 的零值是0
	// 判断key 是否存在
	age, ok := nameAgeMap["ll"]
	if ok {
		fmt.Println(age)
	}
	// 删除
	// delete(nameAgeMap, "rl")
	// fmt.Println(nameAgeMap)

	// 遍历
	for k, v := range nameAgeMap {
		fmt.Printf("数组索引:%s,对应值:%d\n", k, v)
	}
	// 使用一个值的时候只拿出key
	for k := range nameAgeMap {
		fmt.Printf("数组索引:%s\n", k)
	}

	// string 和 []byte, 字符串也是不可变子节序列
	s := "hello 飞雪无情"
	bs := []byte(s)

	fmt.Println(bs)
	// 一个汉子对应的是三个字节
	fmt.Println(s[0], s[1], s[15])
	fmt.Println(len(s))

	fmt.Println(utf8.RuneCountInString(s))

	// for range 会自动的隐式解码unicode字符串,变成9个

	for i, r := range s {

		fmt.Println(i, r)

	}

}
