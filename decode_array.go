package bin

import (
	"reflect"
	"unsafe"

	"github.com/webmafia/fast/binary"
)

func getArrayDecoder(typ reflect.Type, opt *Options) (Decoder, error) {
	elem := typ.Elem()
	itemSize := elem.Size()
	len := elem.Len()
	dec, err := getDecoder(elem, opt)

	if err != nil {
		return nil, err
	}

	return func(r *binary.StreamReader, ptr unsafe.Pointer) (err error) {
		for i := 0; i < len; i++ {
			if err = dec(r, unsafe.Add(ptr, uintptr(i)*itemSize)); err != nil {
				return
			}
		}

		return
	}, nil
}
