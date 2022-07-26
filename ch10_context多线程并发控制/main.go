package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	/*
	context 用来管理多个协程的取消操作：
		四个方法：
			deadline() 方法可以获取设置的戒指时间，返回两个值，第一个值截止的时间，第二个值是否设置了截止时间
			Done()  返回一个只读的channel，如果chan可以读取，就代表了context发起了取消信号
			Err()  返回取消原因，因为什么取消
			Value()  方法获取context绑定的值

		功能上分，四种实现好的context
			空context： 不可取消，没有截止时间，主要用于context树的根结点
			可取消的context-WithCancel: 用于发出取消信号，当取消，子context也会取消
			可定时取消的context-WithDeadline-WithTimeout: 多一个定时功能
			值context-WithValue: 用于存储一个key-value键值对

		使用原则：
			context 不要放在结构体中，以参数方式传递
			context 作为参数，要放在第一位
			要使用context.Background 生成根节点context
			context 传值要传递必须的值，尽可能的少
			context 多协程安全

	 */

	var wg sync.WaitGroup

	wg.Add(4)
	//stopCh := make(chan bool)
	ctx, stop := context.WithCancel(context.Background())
	vCtx := context.WithValue(ctx, "userId", 2)

	go func() {
		defer wg.Done()
		getUser(vCtx)
	}()

	go func() {
		defer wg.Done()
		watchDog(ctx, "1号")
	}()

	go func() {
		defer wg.Done()
		watchDog(ctx, "2号")
	}()

	go func() {
		defer wg.Done()
		watchDog(ctx, "3号")
	}()

	time.Sleep(5 * time.Second)
	stop()
	wg.Wait()
}

func getUser(vCtx context.Context)  {
	for {
		select {
		case <- vCtx.Done():
			fmt.Println("获取用户，退出")
			return
		default:
			userId := vCtx.Value("userId")
			fmt.Println("获取用户，用户id：", userId)
			time.Sleep(time.Second)
		}
	}
}

func watchDog(ctx context.Context, name string)  {

	for {
		select {
		case <- ctx.Done():
			fmt.Println(name, "执行退出")
			return
		default:
			fmt.Println(name, "正在监控。。。")
		}
		time.Sleep(time.Second)
	}

}
