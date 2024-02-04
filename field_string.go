package bin

import (
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

type stringField struct {
	offset uintptr
}

func (f stringField) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	str := *(*string)(ptr)
	b.WriteUvarint(uint64(len(str)))
	b.WriteString(str)
}

func (f stringField) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	n := b.ReadUvarint()
	*(*string)(ptr) = b.ReadString(int(n))
}
