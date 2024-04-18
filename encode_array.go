package bin

import (
	"reflect"
	"unsafe"

	"github.com/webmafia/fast/binary"
)

func getArrayEncoder(typ reflect.Type, opt *Options) (Encoder, error) {
	elem := typ.Elem()
	itemSize := elem.Size()
	len := elem.Len()
	enc, err := getEncoder(elem, opt)

	if err != nil {
		return nil, err
	}

	return func(w *binary.StreamWriter, ptr unsafe.Pointer) (err error) {
		for i := 0; i < len; i++ {
			if err = enc(w, unsafe.Add(ptr, uintptr(i)*itemSize)); err != nil {
				return
			}
		}

		return
	}, nil
}
