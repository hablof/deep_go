package main

import (
	"testing"
)

// go test -bench=. comparison_test.go

const size = 1 << 18

func BenchmarkWithoutOptimizations(b *testing.B) {
	data := make([]int, size)
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(data); j++ {
			data[j] = i
		}
	}
}

func BenchmarkWithCalculatedLength(b *testing.B) {
	data := make([]int, size)
	for i := 0; i < b.N; i++ {
		length := len(data)
		for j := 0; j < length; j++ {
			data[j] = i
		}
	}
}

func BenchmarkWithLoopUnwinding(b *testing.B) {
	data := make([]int, size)
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(data)/4; j += 4 {
			for k := 0; k < 4; k++ {
				data[j+k] = i
			}
		}
	}
}

func BenchmarkWithLoopUnwinding16(b *testing.B) {
	data := make([]int, size)
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(data)/16; j += 16 {
			for k := 0; k < 16; k++ {
				data[j+k] = i
			}
		}
	}
}

func BenchmarkWithLoopUnwinding32(b *testing.B) {
	data := make([]int, size)
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(data)/32; j += 32 {
			for k := 0; k < 32; k++ {
				data[j+k] = i
			}
		}
	}
}
