package bin

import (
	"reflect"
	"unsafe"

	"github.com/webmafia/fast/binary"
)

func getSliceEncoder(typ reflect.Type, opt *Options) (Encoder, error) {
	type sliceHeader struct {
		data unsafe.Pointer
		len  int
		cap  int
	}

	elem := typ.Elem()
	itemSize := elem.Size()
	enc, err := getEncoder(elem, opt)

	if err != nil {
		return nil, err
	}

	return func(w *binary.StreamWriter, ptr unsafe.Pointer) (err error) {
		head := (*sliceHeader)(ptr)
		w.WriteUvarint(uint64(head.len))

		for i := 0; i < head.len; i++ {
			if err = enc(w, unsafe.Add(head.data, uintptr(i)*itemSize)); err != nil {
				return
			}
		}

		return
	}, nil
}
