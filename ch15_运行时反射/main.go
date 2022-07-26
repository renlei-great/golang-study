package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type person struct {
	Name string `json:"name1" bson:"name2"`
	Age int `json:"age1" bson:"age2"`
}

func (p person) Print(a, b string) {
	fmt.Println("输出：", a, b)
}

func (p person) String() string  {
	return fmt.Sprintf("name: %s, age %d", p.Name, p.Age)

}

func main() {
	/*
	反射就是提供一种可以在运行时修改任意类型对象的能力，就比如查看一个字段有哪些方法，有哪些字段，对哪些字段值做修改

	interface{} 是空接口，可以表示任意值

	  在反射的定义中，任何接口都是由类型和值组成，变量的反射表示的就是<type, value>，标准库提供了两种类型，reflect.Value和reflect.Type并且
	还提供了两个函数reflect.ValueOf和reflect.TypeOf分别用来获取变量的type和value
		value结构体中的属性都是私有的，操作值时使用这个
			elem 用来获取指针指向的值
				修改值一般创建反射对象是传入的就是指针，修改时需要先用此方法进行获取原始对象
			interface  用来获取对应的原始类型，转化回原始类型
				iv := reflect.ValueOf(i)
				iv.interface().(int)
			IsNil  值是否为空
			IsZero  值是否为零值
			Kind  方法返回一个kind类型的值，它是一个常量
		type 如何和类型有关的最好用这个，但是这个也可以用于值有关的操作
			Implements 方法用于判断是否实现了接口 u；
			AssignableTo 方法用于判断是否可以赋值给类型 u，其实就是是否可以使用 =，即赋值运算符；
			ConvertibleTo 方法用于判断是否可以转换成类型 u，其实就是是否可以进行类型转换；
			Comparable 方法用于判断该类型是否是可比较的，其实就是是否可以使用关系运算符进行比较

	利用反射实现字符串和结构体的互换，可以使用go语言的标准包 json
		json 转 struct
		struct 转 json
		struct tag 是结构体字段的标记，json转换时会使用tag 的json做key这是 go语言自带json包解析JSON的一种约定

	反射定律
		反射是计算机语言程序检视其自身结构的一种方式，属于原编程，灵活强大，但是存在着不安全，可以绕过很多编译器的静态检查
		1。 任何接口值interface{} 都可以反射出反射对象
		2。 反射对象也可以还原interface{}变量，第一规则可逆
		3。 要修改反射对象，值必须可设置，也就是可寻址

	反射虽然强大，但是过度使用会让代码变得复杂混乱。

	 */
	structToJson()
}

func structToJson(){
	p := person{Name: "rl", Age: 18}
	pt := reflect.TypeOf(p)
	pv := reflect.ValueOf(p)
	num := pt.NumField()

	jsonBuilder := strings.Builder{}
	jsonBuilder.WriteString("{")

	for i:=0; i<pt.NumMethod(); i++{
		//pvm:=pv.Method(i)
		pvm := pv.MethodByName("Print")
		args := []reflect.Value{reflect.ValueOf("登录"),reflect.ValueOf("登录b")}
		pvm.Call(args)
	}
	for i:=0; i<num; i++{

		// 获取tag
		ptf := pt.Field(i)
		tag := ptf.Tag.Get("json")

		// 获取值
		pvf := pv.Field(i)

		// 写
		jsonBuilder.WriteString(fmt.Sprintf("\""+tag+"\": \"%v\"", pvf))

		if i != num-1{
			jsonBuilder.WriteString(",")
		}
	}
	jsonBuilder.WriteString("}")
	fmt.Println(jsonBuilder.String())



}

func test(){

	i := 3
	iv := reflect.ValueOf(i)
	it := reflect.TypeOf(i)
	fmt.Println(iv , it)

	i1 := iv.Interface().(int)
	fmt.Println(i1)

	ipv := reflect.ValueOf(&i)
	ipv.Elem().SetInt(7)
	fmt.Println(i)

	p := person{Name: "rl", Age: 18}
	ppv := reflect.ValueOf(&p)
	fmt.Println(ppv.Kind())
	ppv.Elem().FieldByName("Name").SetString("jsjs")
	fmt.Println(p)
	ppv.Elem().Field(0).SetString("rll")
	fmt.Println(p)

	pv := reflect.ValueOf(p)
	fmt.Println(pv.Kind())


	pt := reflect.TypeOf(p)

	// 遍历字段名
	for i:=0; i<pt.NumField(); i++{
		fmt.Println("字段：", pt.Field(i).Name)
	}

	// 遍历方法
	for i:=0; i<pt.NumMethod(); i++{
		fmt.Println("方法：",pt.Method(i).Name)
	}

	// 判断是否实现了某个接口
	fmt.Println("是否实现了：", pt.(fmt.Stringer))
	_, ok := pt.(io.Writer)
	fmt.Println("是否实现了：", ok)

	// 字符串和结构体进行互换
	ph := person{Name: "huh", Age: 18}
	// struct to json
	jsonB, err := json.Marshal(ph)
	if err == nil {
		fmt.Println("struct to json:", string(jsonB))
	}

	// json to struct
	var convJ person
	respJSON:="{\"name1\":\"李四\",\"age1\":40}"
	json.Unmarshal([]byte(respJSON), &convJ)
	fmt.Println("json to struct:", convJ.Name, "age:", convJ.Age)

	// 通过反射获取tag， 使用field获取到对应的反射字段，此时返回了一个structField，有一个字段是Tag,这个字段存储了所有的tag
	for i:=0; i<pt.NumField();i++{
		sf := pt.Field(i)
		fmt.Println(sf.Tag.Get("json"))
		fmt.Println(sf.Tag.Get("bson"))
	}
	// 结构体可以有多个tag

}
