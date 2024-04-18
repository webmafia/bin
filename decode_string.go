package bin

import (
	"unsafe"

	"github.com/webmafia/fast"
	"github.com/webmafia/fast/binary"
)

func stringDecoder(r *binary.StreamReader, ptr unsafe.Pointer) (err error) {
	n := r.ReadUvarint()
	buf := make([]byte, n)

	if err = r.ReadFull(buf); err != nil {
		return
	}

	*(*string)(ptr) = fast.BytesToString(buf)

	return
}
