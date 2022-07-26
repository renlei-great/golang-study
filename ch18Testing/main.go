package main

func main() {
	/*
	这节主要学习怎么保证代码质量

	单元代码逻辑测试：
		规则：
			1、含有单元测试代码的go文件必须以_test.go结尾
			2、_test.go前面最好是被测试的函数所在的go文件名
			3、单元测试的函数名必须以Test开头
			4、测试函数必须接收一个testing.t 类型的指针，并不反悔任何值
			5、函数名最好是Test + 测试函数名
		命令：
			执行单元测试：go test -v /ch18Testing   此命令会检测所有ch18Testing目录下的单元测试
			查看测试覆盖率： go test -v --coverprofile=ch18.cover ./ch18
			生成一个覆盖率报告：go tool cover -html=ch18.cover -o=ch18.html

	基准测试：是一项用于测试和评估软件性能指标的方法，主要是评估代码性能
		规则：
			1、必须以Benchmark开头
			2、函数必须接收一个testing.B指针，不返回任何值
			3、测试的代码放到for循环中
			4、b.N是基准测试框架提供的，表示循环次数
		命令：
			启动基准测试：go test -bench=. ./ch18
		特殊函数：
			重置计数器：b.ResetTimer()
			开启内存统计：b.ReportAllocs()
		并发基准测试：使用b.RunParallel()
	 */
}

var cache = map[int]int{}

func Fibonacci(n int) int {

	if v, ok := cache[n]; ok{
		return v
	}

	var result int
	switch {
	case n <0:
		result = 0
	case n == 0:
		result = 0
	case n == 1:
		result = 1
	default:
		result = Fibonacci(n-1) + Fibonacci(n-2)
	}
	cache[n] = result
	return result
}
