package bin

import (
	"unsafe"

	fast "github.com/webmafia/fast"
)

var _ Type = raw{}

// TODO? Use this if architecture is little-endian and the type contains no pointers. If the type is a struct, it can't have any unexported fields.
type raw struct {
	size int
}

func (r raw) bytes(ptr unsafe.Pointer) []byte {
	header := sliceHeader{
		data: ptr,
		len:  r.size,
		cap:  r.size,
	}

	return *(*[]byte)(unsafe.Pointer(&header))
}

// decode implements Type.
func (r raw) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader, _ bool) (err error) {
	_, err = b.Read(r.bytes(ptr))
	return
}

// encode implements Type.
func (r raw) encode(ptr unsafe.Pointer, b fast.Writer) {
	b.Write(r.bytes(ptr))
}

// encodedSize implements Type.
func (r raw) encodedSize(ptr unsafe.Pointer) int {
	return r.size
}
