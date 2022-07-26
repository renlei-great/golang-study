package main

import (
	"errors"
	"fmt"
)

type person struct {
	name string
	age uint
}

type WalkRun interface {
	Walk()
	Run()
}

type commonError struct {
	errorCode int
	errorMsg string
}

type MyError struct {
	err error
	msg string
}

func (e *MyError) Error() string{
	return e.err.Error() + e.msg
}

func (ce *commonError) Error() string {
	return ce.errorMsg
	//return "为啥呢"
}

func (p *person) Walk() {
	fmt.Printf("%s能走\n", p.name)
}

func (p *person) Run() {
	fmt.Printf("%s能跑\n", p.name)
}

func add(a,b int) (int, error) {
	if a < 0 || b < 0{
		return 0, &commonError{errorCode: 100, errorMsg: "a或者b不能为负数"}
	} else {
		return a+b, nil
	}
}

func connectMySQL(ip, username, password string){
	if ip == ""{
		panic("ip不能为空")
	}
}

func main(){
	//i, err := strconv.Atoi("a")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(i)
	//}

	defer func() {
		if p:=recover(); p!=nil{
			fmt.Println(p)
		}
	}()
	fmt.Println(111)
	connectMySQL("", "root", "123")
	fmt.Println(222)


	sum, err := add(-1, 2)

	if cm, ok := err.(*commonError); ok {
		fmt.Println("错误代码为:", cm.errorCode, ", 错误信息为：", cm.errorMsg)
	} else {
		fmt.Println(sum)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}

	//newErr := MyError{err, "数据上传问题"}

	w := fmt.Errorf("新的错误是:%w", err)
	fmt.Println(w)
	fmt.Println(errors.Unwrap(w))
	fmt.Println(errors.Is(w, err))

	var cm *commonError
	if errors.As(w, &cm){
		fmt.Println("错误代码为：", cm.errorCode, "，错误形象为", cm.errorMsg)
	}else {
		fmt.Println(sum)
	}
}
