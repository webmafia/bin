package bin

import (
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

type stringType struct {
	offset uintptr
}

func (f stringType) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	str := *(*string)(unsafe.Add(ptr, f.offset))
	b.WriteUvarint(uint64(len(str)))
	b.WriteString(str)
}

func (f stringType) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) error {
	n := b.ReadUvarint()
	*(*string)(unsafe.Add(ptr, f.offset)) = b.ReadString(int(n))
	return nil
}
