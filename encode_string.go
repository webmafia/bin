package bin

import (
	"unsafe"

	"github.com/webmafia/fast/binary"
)

var _ Encoder = stringEncoder

func stringEncoder(w *binary.StreamWriter, ptr unsafe.Pointer) error {
	str := *(*string)(ptr)
	w.WriteUvarint(uint64(len(str)))
	w.WriteString(str)
	return nil
}
