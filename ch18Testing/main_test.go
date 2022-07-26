package main

import "testing"

// 代码逻辑测试
func TestFibonacci(t *testing.T) {
	fsMap := map[int]int{}

	fsMap[-1] = 0
	fsMap[0] = 0
	fsMap[1] = 1
	fsMap[2] = 1
	fsMap[3] = 2
	fsMap[4] = 3
	fsMap[5] = 5
	fsMap[6] = 8
	fsMap[7] = 13
	fsMap[8] = 21
	fsMap[9] = 34

	for k, v := range fsMap{
		fr := Fibonacci(k)
		if v == fr{
			t.Logf("结果正确n为%d, 值为%d", k, fr)
		} else{
			t.Errorf("结果错误n为%d, 值为%d", k, fr)
		}
	}
}

// 基准测试
func BenchmarkFibonacci(b *testing.B) {
	n := 10

	b.ReportAllocs()  // 开启内存统计
	b.ResetTimer()  // 重置计时器
	for i := 0; i < b.N; i++ {
		Fibonacci(n)
	}
}

// 并发测试
func BenchmarkFibonacciRunParalle(b *testing.B) {
	n:=10

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			Fibonacci(n)
		}
	})
}