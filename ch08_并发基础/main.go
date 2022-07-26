package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main1(){
	// 简单使用sleep进行阻塞实现协程并发
	go fmt.Println("我是rl1")
	fmt.Println("我是main goroutine打印的")
	time.Sleep(time.Second)

}

func main2() {
	// 使用channel协程进行通讯
	ch := make(chan string)

	go func() {
		fmt.Println("我是rl")
		ch <- "goroutine 完成"
	}()

	fmt.Println("我是 main goroutine ")
	v:= <- ch
	fmt.Println("接收到的chan的值为:v", v)

}

func counter(out chan<- int) {
	out <- 2
	out <- 2

}

func jter(j <-chan int ) {
	fmt.Println(<-j)
	fmt.Println(<-j)

}

func main() {

	/*
		使用go 关键字启动一个goroutine, 所有的goroutine都由 go runtime 所调度，而go语言中没有线程的概念，go语言中的并发是由go自己进行
		调度的。使用go关键字启动一个goroutine

		go语言中有一种channel类型，用于携程之间的通信。
			无缓存的channel，容量是0，不能存储任何数据，只能起到传输数据作用，也称为同步channel
			有缓存channel,是一个类似可阻塞的队列，先进先出。
			可以使用cap和len，进行容量和长度的获取
			可以使用close(chan)关闭， 关闭后就不可以往里写数据了，但是还能读，如果读的没有了，返回元素的零值
			*还有单项channel,在函数的参数类型中指名，如果违反会编译错误
			** 数据流动、传递的场景中要优先使用channel，它是并发安全的，性能也好

		select多路复用实现：
			类似switch语法，但是它的case后只跟channel，如果获取不到值会一直阻塞


	 */

	cacheCh := make(chan int, 5)
	cacheCh <- 2
	cacheCh <- 3
	fmt.Println(cap(cacheCh), len(cacheCh))
	counter(cacheCh)
	jter(cacheCh)

	firstCh := make(chan string)
	secondCh := make(chan string)
	threeCh := make(chan string)



	go func() {
		firstCh <- downloadFile("firstCh")
	}()
	go func() {
		secondCh <- downloadFile("secondCh")
	}()
	go func() {
		threeCh <- downloadFile("threeCh")
	}()

	// 使用select 如果某个通道有值，就执行其中的程序体，如果那个都执行不了将一直等待
	select {
		case filePath := <- firstCh:
			fmt.Println(filePath)

		case filePath := <- secondCh:
			fmt.Println(filePath)

		case filePath := <- threeCh:
			fmt.Println(filePath)
	}


}

func downloadFile(chanName string) string{
	rand.Seed(time.Now().Unix())
	s := rand.Intn(5)
	time.Sleep(time.Second * time.Duration(s))
	return chanName + ":filePath"
}

func moreDefer(){
	// 上节课defer复习
	defer fmt.Printf("1111")
	defer fmt.Printf("2222")
	defer fmt.Printf("3333")

}

