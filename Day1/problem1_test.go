package day1

import "testing"

func BenchmarkProblem1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem1()
	}
}
