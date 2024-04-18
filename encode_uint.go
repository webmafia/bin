package bin

import (
	"unsafe"

	"github.com/webmafia/fast/binary"
)

func uintEncoder(w *binary.StreamWriter, ptr unsafe.Pointer) error {
	return w.WriteUint(*(*uint)(ptr))
}

func uint8Encoder(w *binary.StreamWriter, ptr unsafe.Pointer) error {
	return w.WriteUint8(*(*uint8)(ptr))
}

func uint16Encoder(w *binary.StreamWriter, ptr unsafe.Pointer) error {
	return w.WriteUint16(*(*uint16)(ptr))
}

func uint32Encoder(w *binary.StreamWriter, ptr unsafe.Pointer) error {
	return w.WriteUint32(*(*uint32)(ptr))
}

func uint64Encoder(w *binary.StreamWriter, ptr unsafe.Pointer) error {
	return w.WriteUint64(*(*uint64)(ptr))
}
