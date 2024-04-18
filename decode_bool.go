package bin

import (
	"unsafe"

	"github.com/webmafia/fast/binary"
)

func boolDecoder(r *binary.StreamReader, ptr unsafe.Pointer) error {
	*(*bool)(ptr) = r.ReadBool()
	return r.Error()
}
