package bin

import (
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

type boolType struct {
	offset uintptr
}

func (f boolType) encodedSize(_ unsafe.Pointer) int {
	return 1
}

func (f boolType) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteBool(*(*bool)(unsafe.Add(ptr, f.offset)))
}

func (f boolType) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) error {
	*(*bool)(unsafe.Add(ptr, f.offset)) = b.ReadBool()
	return nil
}
