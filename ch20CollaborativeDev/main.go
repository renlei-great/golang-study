package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gotour/ch20CollaborativeDev/util"
)

func main() {
	/*
	包：
		通过package定义一个包，一个包就是一个独立的空间，包里定义了结构体，函数等
		使用包使用import导入一个包
		包中的首字母是否大写表示了他的作用域是包内还是包外
		init函数导包时就会执行，在main函数之前执行，init函数可以有多个，但是执行顺序时随机的

	模块：
		创建模块或项目：go mod init host/name
			模块名最好以域名开头
		使用第三方模块：
			go env -w GO111MODULE=on
			go env -w GOPROXY=https://ggoproxy.io,direct
			设置不走 proxy 的私有仓库: go env -w GOPRIVATE=*.corp.example.com
			go get -u github.com/gin-gonic/gin
			同步go.mod文件中确实的模块 go mod tidy

	 */

	r := gin.Default()
	r.Run()
	fmt.Println("我是main")

	util.Test()

}
