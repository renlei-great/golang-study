package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum int
	mut sync.RWMutex

)

func main() {
	/*
	这节主要是使用sync包
		1。 使用sync.WaitGroup控制等待协程完成进行程序的急时退出
			Add(num)  添加有多少需要等待的协程
			Done()   一个协程执行完成，需要等待的协程数减1，此方法可以配合defer使用
			Wait()   阻塞等待所有的协程

		2。 互斥锁 sync.mutex 在出现资源抢夺的时候使用互斥锁实现同一时刻只能有一个协程在执行某段代码
			Lock()  上锁
			Unlock()  释放锁， 也可配合defer使用
			这两个方法往往是成对出现

		3。 读写锁，可以通过sync.RWMutex 提升加锁后的性能，通过读可以同时读，提升性能
			写锁 Lock()
			释放写锁 Unlock()
			读锁 RLock()
			释放读锁 RUnlock()

		4. sync.Once只让代码执行一次，哪怕是高并发的情况下，比如创建一个单例
			var once sync.Once
			Do()  添加要执行的方法，此方法在所有的协程中只会执行一次

		5。 sync.Cond 可以做到同时阻塞多个协程或者同时唤醒多个协程，就是可以做到阻塞或唤醒协程,是基于互斥锁实现的
			var cond sync.NewCond(&sync.Mutex)
			cond.L.Lock()  想要阻塞必须先加锁
			cond.Wait()  阻塞协程
			cond.L.Unlock()  释放锁
			cond.Broadcast()  唤醒所有协程
			Signal  唤醒一个等待时间最长的协程

	 */


	MyMap()
}

func MyMap() {
	var sMap sync.Map

	sMap.Store("rl", 18)
	sMap.Store("ll", 19)
	sMap.Store("gq", 17)

	fmt.Println(sMap.Load("rl"))
	fmt.Println(sMap.Load("rll"))
	fmt.Println(sMap.LoadOrStore("rll", 188))
	fmt.Println(sMap.Load("rll"))


	sMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

}

func rece()  {
	// 开启10个协程，使用sync.Cond 来控制所有的协程同时执行

	cond := sync.NewCond(&sync.Mutex{})  // 生成一个*sync.Cond 可以阻塞和唤醒协程

	var wg sync.WaitGroup  // 创建一个阻塞协程的锁
	wg.Add(11)  // 添加11个协程阻塞

	for i:=0; i<10; i++{
		go func(num int) {
			defer wg.Done()

			fmt.Println(num, "号已经就位")
			cond.L.Lock()  // 上锁
			cond.Wait()  // 阻塞
			fmt.Println(num, "号开始跑")
			cond.L.Unlock()  // 释放锁
		}(i)
	}

	time.Sleep(2 * time.Second)

	go func() {
		defer wg.Done()

		fmt.Println("裁判就位，准备")
		fmt.Println("比赛开始")
		cond.Broadcast()  // 执行所有使用cond阻塞的协程
	}()

	wg.Wait()

}

func doOnce() {
	var once sync.Once
	var wg sync.WaitGroup
	onceBody := func() {
		fmt.Println("only once")
	}

	// 用于等待协程执行完毕
	wg.Add(10)
	//done := make(chan bool)

	// 启动10个协程执行once.Do(onceBody)
	for i:=0; i<10; i++{
		go func() {
			// 把要执行的函数(方法) 作为参数传给once.Do方法即可
			once.Do(onceBody)
			//done <- true
			defer wg.Done()
		}()

	}

	//for i:=0; i<10; i++{
	//	<-done
	//}
	wg.Wait()
}


func run(){

	// 监控协程的
	var wg sync.WaitGroup

	// 因为要监控110个协程
	wg.Add(110)

	for i := 0; i < 100; i++{
		go func() {
			defer wg.Done()
			add(10)
		}()
	}

	for i:=0; i<10;i++{
		go func() {
			defer wg.Done()
			fmt.Println("和为：", readSum())
		}()

	}

	// 一直等待，直到计数器为0
	wg.Wait()
}

func add(i int) {
	mut.Lock()
	defer mut.Unlock()
	sum += i
	//mut.Unlock()
}

func readSum() int {
	mut.RLock()
	defer mut.RUnlock()
	b := sum
	return b
}
