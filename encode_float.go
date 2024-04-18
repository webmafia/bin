package bin

import (
	"unsafe"

	"github.com/webmafia/fast/binary"
)

func float32Encoder(w *binary.StreamWriter, ptr unsafe.Pointer) error {
	return w.WriteFloat32(*(*float32)(ptr))
}

func float64Encoder(w *binary.StreamWriter, ptr unsafe.Pointer) error {
	return w.WriteFloat64(*(*float64)(ptr))
}
