package fib

import "testing"

// Fib 基准测试

func BenchMarkFib(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Fib(2)
	}
}

// 内部调用的函数
func benchMarkFib(b *testing.B, n int)  {
	for i := 0; i < b.N; i++ {
		Fib(2)
	}
}

func BenchMarkFib2(b *testing.B)  {
	benchMarkFib(b, 2)
}

func BenchMarkFib20(b *testing.B)  {
	benchMarkFib(b, 20)
}

func BenchMarkFib200(b *testing.B)  {
	benchMarkFib(b, 200)
}

// 并行测试

func BenchmarkSplitParallel(b *testing.B)  {
	// b.SetParallelism(1) // 设置使用CPU的数量
	b.RunParallel(func (pb *testing.PB)  {
		for pb.Next() {
			Split("ancna", "c")
		}
	})
}