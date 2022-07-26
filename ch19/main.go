package main

import (
	"fmt"
	"os"
)

const name = "rlei"

func main() {
	/*
	代码规范检查，静态扫描检查：
		包：golangci-lint
		相关配置：.golangci.yml
		写正确的代码是性能优化的前提

	堆分配还是栈分配，go语言两部分内存空间就是堆内存和栈内存
		栈内存由编译器自动分配和释放，开发者无法控制，一般存放一些局部变量，参数等，函数创建的时候，这些内存自动创建，返回的时候自动释放
		堆内存的生命周期会比栈长，如果函数返回的值在其他地方还有使用，那么这个值就会被编译器分配给堆上，堆内存只能被垃圾回收器释放，所以栈内存的效率要高
		逃逸分析：
			可以查看被分配到了堆内存还是栈内存
			命令：go build -gcflags="-m -l" ./ch19/main.go
			常见逃逸：
				1、指针做为函数返回值的时候，肯定会逃逸
				2、被要已经逃逸的指针引用的变量也会发送逃逸

	优化技巧：
		1、尽可能的避免逃逸，因为栈的效率更高，
		2、如果避免不了逃逸，那么要尽可能的重用内存，比如使用sync.pool
		3、使用合适的算法

	解析性能的工具 pprof，可以看到cpu,内存，阻塞，互斥锁

	 */

	os.Mkdir("tmp",0666)
	newString()
	fmt.Println("rlll")

	m := map[string]*string{}
	s := "rrl"
	m["1"] = &s


}

func newString() *string{
	s := new(string)
	*s = "relelel"
	return s
}
