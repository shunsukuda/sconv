package sconv

import (
	"strconv"
	"testing"
)

func BenchmarkStr19DecTouint64(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Str19DecToU64("1234567890123456789")
	}
	b.StopTimer()
}

func BenchmarkStrconvAtoi19(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = strconv.Atoi("1234567890123456789")
	}
	b.StopTimer()
}

func BenchmarkStrconvParse19(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseUint("1234567890123456789", 10, 64)
	}
	b.StopTimer()
}
