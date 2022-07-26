package main

import "fmt"

type address struct {
	province string
	city string
}

type person1 struct {
	name string
	age uint
	city string
	address
}

type person struct {
	name string
	age uint
}

type age uint

type Stringer interface {
	String() string
}

func (p person) String() string {
	return fmt.Sprintf("the name is %s, age is %d", p.name ,p.age)
	//return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
}

func (age age) String() string {
	return fmt.Sprintf("the name is %d, age is %d", age ,age)
	//return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
}

func (addr address) String() string {
	return fmt.Sprintf("the province is %s, city is %s", addr.province ,addr.city)
	//return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
}

func prinString(s Stringer){
	fmt.Println(s.String())
}

func NewPerson(name string) *person {
	// 工厂函数
	return &person{name: name}
}

func New(text string) error {
	return &errorString{text}
}

// 接口体，内部s字段存储错误信息
type errorString struct {
	s string
}

// 用于实现error接口
func (e *errorString) Error() string {
	return e.s
}




func main(){
	//var p person
	//p:=person{age:30,name:"飞雪无情"}
	p := person{"renlei", 19}
	fmt.Println(p)
	fmt.Println(p.String())
	prinString(p)
	a := age(15)
	prinString(a)
	prinString(&a)

	p1 := NewPerson("ll")
	fmt.Println(p1.name, p1.age)

	fmt.Println(New("111222333").Error())

	p11 := person1{
		name: "renlei",
		age:  18,
		city: "商都",
		address: address{
			province: "qqq",
			city:     "北京",
		},
	}
	fmt.Println(p11.city)
	fmt.Println(p11.province)
	prinString(p11)
	fmt.Println("断言")
	var s fmt.Stringer
	s = p1
	p2:=s.(*person)
	fmt.Println(p2)
	//addr := address{province: "aa", city: "bb"}
	aa, ok :=s.(address)
	if ok {
		fmt.Println(aa)
	} else {
		fmt.Println("s不是一个address")
	}



}
