package __benchmark

import "testing"

// 基准测试 go test -bench .

// 串行基准测试
func BenchmarkSelectServer(b *testing.B) {
	InitServerIndex()
	b.ResetTimer() // 重置时间 因为不计初始化时间
	for i := 0; i < b.N; i++ {
		SelectServer()
	}
}

// 并行基准测试
func BenchmarkSelectServerParallel(b *testing.B) {
	InitServerIndex()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SelectServer()
		}
	})
}

// 使用内置rand包，因为内置rand包会在生成随机数时全局上锁
// king@chengxiaofengdeiMac 5_benchmark % go test -bench .
// goos: darwin
// goarch: amd64
// pkg: Go_UnitTest/5_benchmark
// cpu: Intel(R) Core(TM) i5-4670T CPU @ 2.30GHz
// BenchmarkSelectServer-4                 49767816                22.07 ns/op
// BenchmarkSelectServerParallel-4          6652312               159.3 ns/op
// PASS
// ok      Go_UnitTest/5_benchmark 2.376s

// 使用fastrand包，速度明显加快
// king@chengxiaofengdeiMac 5_benchmark % go test -bench .
// goos: darwin
// goarch: amd64
// pkg: Go_UnitTest/5_benchmark
// cpu: Intel(R) Core(TM) i5-4670T CPU @ 2.30GHz
// BenchmarkSelectServer-4                 216198817                5.663 ns/op
// BenchmarkSelectServerParallel-4         703860760                1.636 ns/op
// PASS
// ok      Go_UnitTest/5_benchmark 3.123s
