package hash_benchmark

import (
	"crypto"
	"fmt"
	"log"
	"testing"

	_ "golang.org/x/crypto/blake2b"
	_ "golang.org/x/crypto/blake2s"
	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
	_ "golang.org/x/crypto/sha3"

	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
)

func genUUID() []byte {
	b := make([]byte, 16)
	_, err := rnd.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return []byte(fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]))
}

var uuidStr = genUUID()

func BenchmarkMD4(b *testing.B) {
	h := crypto.MD4.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkMD5(b *testing.B) {
	h := crypto.MD5.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkSHA1(b *testing.B) {
	h := crypto.SHA1.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkSHA224(b *testing.B) {
	h := crypto.SHA224.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkSHA256(b *testing.B) {
	h := crypto.SHA256.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkSHA384(b *testing.B) {
	h := crypto.SHA384.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkSHA512(b *testing.B) {
	h := crypto.SHA512.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkSHA512_224(b *testing.B) {
	h := crypto.SHA512_224.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkSHA512_256(b *testing.B) {
	h := crypto.SHA512_256.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkRIPEMD160(b *testing.B) {
	h := crypto.RIPEMD160.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkSHA3_224(b *testing.B) {
	h := crypto.SHA3_224.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkSHA3_256(b *testing.B) {
	h := crypto.SHA3_256.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkSHA3_384(b *testing.B) {
	h := crypto.SHA3_384.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkSHA3_512(b *testing.B) {
	h := crypto.SHA3_512.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkBLAKE2s_256(b *testing.B) {
	h := crypto.BLAKE2s_256.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkBLAKE2b_256(b *testing.B) {
	h := crypto.BLAKE2b_256.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkBLAKE2b_384(b *testing.B) {
	h := crypto.BLAKE2b_384.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}

func BenchmarkBLAKE2b_512(b *testing.B) {
	h := crypto.BLAKE2b_512.New()
	for i := 0; i < b.N; i++ {
		h.Sum(uuidStr)
	}
}
