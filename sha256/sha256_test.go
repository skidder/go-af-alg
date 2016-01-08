package sha256

import (
	"crypto/rand"
	"crypto/sha256"
	"io"
	"testing"
)

func TestSHA256(t *testing.T) {
	for i := 0; i < 100; i++ {
		buf := make([]byte, 10*1024-i)
		if _, err := io.ReadFull(rand.Reader, buf); err != nil {
			t.Fatal(err)
		}

		expected := sha256.Sum256(buf)
		got, err := SHA256(buf)
		if err != nil {
			t.Fatal(err)
		}

		if expected != got {
			t.Fatal("exp:%x got:%x", expected, got)
		}
	}
}

func TestSHA256Writer(t *testing.T) {
	hash := sha256.New()

	for i := 0; i < 100; i++ {
		ohash, err := New()
		if err != nil {
			t.Fatal(err)
		}
		hash.Reset()
		buf := make([]byte, 10*1024-i)
		if _, err := io.ReadFull(rand.Reader, buf); err != nil {
			t.Fatal(err)
		}

		if _, err := ohash.Write(buf); err != nil {
			t.Fatal(err)
		}
		if _, err := hash.Write(buf); err != nil {
			t.Fatal(err)
		}

		var got, exp [32]byte

		hash.Sum(exp[:0])
		got, err = ohash.Sum()
		if err != nil {
			t.Fatal(err)
		}

		if got != exp {
			t.Fatal("exp:%x got:%x", exp, got)
		}
	}
}

type shafunc func([]byte)

func benchmarkSHA256(b *testing.B, length int64, fn shafunc) {
	buf := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		b.Fatal(err)
	}
	b.SetBytes(length)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(buf)
	}
}

func BenchmarkSHA256Large(b *testing.B) {
	benchmarkSHA256(b, 1024*1024, func(buf []byte) { SHA256(buf) })
}

func BenchmarkSHA256Large_stdlib(b *testing.B) {
	benchmarkSHA256(b, 1024*1024, func(buf []byte) { sha256.Sum256(buf) })
}

func BenchmarkSHA256Small(b *testing.B) {
	benchmarkSHA256(b, 1, func(buf []byte) { SHA256(buf) })
}

func BenchmarkSHA256Small_stdlib(b *testing.B) {
	benchmarkSHA256(b, 1, func(buf []byte) { sha256.Sum256(buf) })
}
