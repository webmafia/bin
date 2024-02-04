package bin

import (
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

type field interface {
	encode(ptr unsafe.Pointer, b *fast.BinaryBuffer)
	decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader)
}
