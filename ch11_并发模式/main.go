package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func buy(n int) <-chan string  {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++{
			out <- fmt.Sprint("配件", i)
		}
	}()

	return out
}

func build(in <- chan string) <-chan string{
	out := make(chan string)

	go func() {
		defer close(out)

		for c:= range in {
			out <- "组装（" + c + ")"
		}
	}()
	return out
}

func pack(in <-chan string) <- chan string{
	out := make(chan string)

	go func() {
		defer close(out)

		for c := range in{
			out <- "打包(" + c + ")"
		}
	}()

	return out
}

func merge(ins ...<-chan string) <-chan string{

	var wg sync.WaitGroup
	out := make(chan string)

	// 把一个chan的数据发送到out中
	p := func(in <-chan string) {
		defer wg.Done()
			for c := range in{
			out <- c
		}
	}

	wg.Add(len(ins))

	// 扇入，需要启动多个goroutine用于处理多个channel中的数据
	for _, cs := range ins{
		go p(cs)
	}

	// 等待所有输入的数据ins处理完，再关闭输出out
	go func() {
		wg.Wait()
		close(out)
	}()
	return out


}

func main() {

	/*
	本节主要学习了并发模式，主要感觉是针对一些场景的并发解决方法
		for select 循坏模式
			非常常见的模式，是一种多路复用模式，那个满足就执行那个，直到满足退出的条件进行退出
		select timeout 模式
			主要是防止select 出现的无限等待，可以使用time.After设置一个超时时间退出时间，但是如果可以优先使用context进行超时取消控制
		pipeline 模式
			流水线模式，有好几道工序，但是每道工序都需要上一道工序的输出，下一道工序等待此刻工序的输出结果作为输入
			也就是说，对于此刻的工序，上一道工序是生产者，下一道工序是消费者
			构成：
				1。 流水线由一道一道的工序构成，每道工序通过channel把数据传到下一工序
				2。 每道工序一般对应着一个函数，函数有协程和channel，协程一般处理数据放入道channel中，整个函数返回的channel供下个函数用
				3。 最终有一个组织者（示例的main函数），把所有工序串起来，变成了一个流水线
		扇出和扇入模式
			是基于流水线进行的并发改进，通过对某一工序增加并发数量实现，再通过新的merge方法进行管道数据的整合使得增加并发的工序对上一道工序和下一道工序还和从前一样
			此代码中的示例还实现了开闭原则规定"软件中的对象（类，模块，函数等等）应该啊对于扩展是开放的，但对于修改是封闭的"
		futures 模式
			未来模式，主协程不用等待子协程返回的结果，可以做其他的事情，等未来需要子协程结果的时候再取，如果没有返回就会一直等待


	 */


	//ctx, stop := context.WithTimeout(context.Background(), 5 * time.Second)
	//timeOut(ctx)
	//
	//stop()

	//coms := buy(100)
	//phones1 := build(coms)
	//phones2 := build(coms)
	//phones3 := build(coms)
	//phones := merge(phones1, phones2, phones3)
	//packs := pack(phones)
	//
	//for p := range packs{
	//	fmt.Println(p)
	//}

	washCh := washVegetables()
	waterCh := boilWater()

	fmt.Println("安排了洗菜和烧水，做别的事")
	time.Sleep(time.Second * 2)
	fmt.Println("看看水和菜好嘞没有")

	wash := <-washCh
	water := <-waterCh

	fmt.Println("开始做火锅", wash, water)

}

func washVegetables() <-chan string{
	vegetables := make(chan string)

	go func() {
		time.Sleep(time.Second * 5)
		vegetables <- "洗好的菜"
	}()

	return vegetables
}

func boilWater() <-chan string{
	water := make(chan string)

	go func() {
		time.Sleep(time.Second * 5)
		water <- "烧好的水"
	}()

	return water
}



func timeOut(ctx context.Context){
	result := make(chan string)

	go func() {
		time.Sleep(8 * time.Second)
		result <- "服务端结果"
	}()

	select {
	case v := <-result:
		fmt.Println(v)
	case <- ctx.Done():
		fmt.Println("ctx访问超时")
	case <- time.After(5 * time.Second):
		fmt.Println("访问超时")
	}
}
