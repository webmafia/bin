package bin

import (
	"testing"
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

func BenchmarkEncode_String(b *testing.B) {
	buf := fast.NewBinaryBuffer(1024)
	enc := stringType{}
	str := "hello there"

	for i := 0; i < b.N; i++ {
		enc.encode(unsafe.Pointer(&str), buf)
		buf.Reset()
	}
}

func BenchmarkEncodeInterface_String(b *testing.B) {
	buf := fast.NewBinaryBuffer(1024)
	var enc Type = stringType{}
	str := "hello there"

	for i := 0; i < b.N; i++ {
		enc.encode(unsafe.Pointer(&str), buf)
		buf.Reset()
	}
}
