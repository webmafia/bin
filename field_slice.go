package bin

import (
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

type sliceField struct {
	offset uintptr
}

func (f sliceField) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	// TODO
}
