package bin

import (
	"unsafe"

	"github.com/webmafia/fast/binary"
)

func float32Decoder(r *binary.StreamReader, ptr unsafe.Pointer) error {
	*(*float32)(ptr) = r.ReadFloat32()
	return r.Error()
}

func float64Decoder(r *binary.StreamReader, ptr unsafe.Pointer) error {
	*(*float64)(ptr) = r.ReadFloat64()
	return r.Error()
}
