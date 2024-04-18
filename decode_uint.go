package bin

import (
	"unsafe"

	"github.com/webmafia/fast/binary"
)

func uintDecoder(r *binary.StreamReader, ptr unsafe.Pointer) error {
	*(*uint)(ptr) = r.ReadUint()
	return r.Error()
}

func uint8Decoder(r *binary.StreamReader, ptr unsafe.Pointer) error {
	*(*uint8)(ptr) = r.ReadUint8()
	return r.Error()
}

func uint16Decoder(r *binary.StreamReader, ptr unsafe.Pointer) error {
	*(*uint16)(ptr) = r.ReadUint16()
	return r.Error()
}

func uint32Decoder(r *binary.StreamReader, ptr unsafe.Pointer) error {
	*(*uint32)(ptr) = r.ReadUint32()
	return r.Error()
}

func uint64Decoder(r *binary.StreamReader, ptr unsafe.Pointer) error {
	*(*uint64)(ptr) = r.ReadUint64()
	return r.Error()
}
