package bin

import (
	"unsafe"

	"github.com/webmafia/fast/binary"
)

type float32Type struct {
	offset uintptr
}

func (f float32Type) encodedSize(_ unsafe.Pointer) int {
	return 4
}

func (f float32Type) encode(ptr unsafe.Pointer, b binary.Writer) {
	b.WriteFloat32(*(*float32)(unsafe.Add(ptr, f.offset)))
}

func (f float32Type) decode(ptr unsafe.Pointer, b binary.Reader, _ bool) error {
	*(*float32)(unsafe.Add(ptr, f.offset)) = b.ReadFloat32()
	return nil
}

type float64Type struct {
	offset uintptr
}

func (f float64Type) encodedSize(_ unsafe.Pointer) int {
	return 4
}

func (f float64Type) encode(ptr unsafe.Pointer, b binary.Writer) {
	b.WriteFloat64(*(*float64)(unsafe.Add(ptr, f.offset)))
}

func (f float64Type) decode(ptr unsafe.Pointer, b binary.Reader, _ bool) error {
	*(*float64)(unsafe.Add(ptr, f.offset)) = b.ReadFloat64()
	return nil
}
