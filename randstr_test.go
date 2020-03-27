package hash_benchmark

import (
	"math/rand"
	"testing"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var rnd *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rnd.Intn(len(charset))]
	}
	return string(b)
}

func BenchmarkRandString8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringWithCharset(8, charset)
	}
}

func BenchmarkRandString16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringWithCharset(16, charset)
	}
}

func BenchmarkRandString32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringWithCharset(32, charset)
	}
}

func BenchmarkRandString64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringWithCharset(64, charset)
	}
}

func BenchmarkRandString128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringWithCharset(128, charset)
	}
}
