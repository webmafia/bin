package bin

import (
	"unsafe"

	"github.com/webmafia/fast/binary"
)

func boolEncoder(w *binary.StreamWriter, ptr unsafe.Pointer) error {
	return w.WriteBool(*(*bool)(ptr))
}
