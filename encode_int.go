package bin

import (
	"unsafe"

	"github.com/webmafia/fast/binary"
)

func intEncoder(w *binary.StreamWriter, ptr unsafe.Pointer) error {
	return w.WriteInt(*(*int)(ptr))
}

func int8Encoder(w *binary.StreamWriter, ptr unsafe.Pointer) error {
	return w.WriteInt8(*(*int8)(ptr))
}

func int16Encoder(w *binary.StreamWriter, ptr unsafe.Pointer) error {
	return w.WriteInt16(*(*int16)(ptr))
}

func int32Encoder(w *binary.StreamWriter, ptr unsafe.Pointer) error {
	return w.WriteInt32(*(*int32)(ptr))
}

func int64Encoder(w *binary.StreamWriter, ptr unsafe.Pointer) error {
	return w.WriteInt64(*(*int64)(ptr))
}
