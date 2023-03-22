package base

import "testing"

func BenchmarkDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(7, 8)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() // 调用该函数停止压力测试的时间计数

	// 做一些初始化，例如读取文件数据，数据库连接之类的
	// 这样这些时间不影响我们测试函数本身的性能

	// b.N = 100000   // 可以修改执行次数
	b.StartTimer() // 重新开始时间计时
	for i := 0; i < b.N; i++ {
		Division(7, 8)
	}
}
