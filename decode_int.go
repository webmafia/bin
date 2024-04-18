package bin

import (
	"unsafe"

	"github.com/webmafia/fast/binary"
)

func intDecoder(r *binary.StreamReader, ptr unsafe.Pointer) error {
	*(*int)(ptr) = r.ReadInt()
	return r.Error()
}

func int8Decoder(r *binary.StreamReader, ptr unsafe.Pointer) error {
	*(*int8)(ptr) = r.ReadInt8()
	return r.Error()
}

func int16Decoder(r *binary.StreamReader, ptr unsafe.Pointer) error {
	*(*int16)(ptr) = r.ReadInt16()
	return r.Error()
}

func int32Decoder(r *binary.StreamReader, ptr unsafe.Pointer) error {
	*(*int32)(ptr) = r.ReadInt32()
	return r.Error()
}

func int64Decoder(r *binary.StreamReader, ptr unsafe.Pointer) error {
	*(*int64)(ptr) = r.ReadInt64()
	return r.Error()
}
