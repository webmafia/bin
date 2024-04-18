package bin

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/webmafia/fast/binary"
)

func getSliceDecoder(typ reflect.Type, opt *Options) (Decoder, error) {
	type sliceHeader struct {
		data unsafe.Pointer
		len  int
		cap  int
	}

	elem := typ.Elem()
	itemSize := elem.Size()
	dec, err := getDecoder(elem, opt)

	if err != nil {
		return nil, err
	}

	if opt.AllowAllocations {
		return func(r *binary.StreamReader, ptr unsafe.Pointer) (err error) {
			head := (*sliceHeader)(ptr)
			calcSize := int(r.ReadUvarint())

			if calcSize > head.cap {

				// Allocate new slice with the calculated size
				newSlice := reflect.MakeSlice(typ, calcSize, calcSize)
				head.data = newSlice.UnsafePointer()
				head.cap = calcSize
				head.len = 0
			}

			for i := 0; i < calcSize; i++ {
				elemPtr := unsafe.Add(head.data, uintptr(i)*itemSize)

				if err = dec(r, elemPtr); err != nil {
					return err
				}
			}

			head.len = calcSize

			return
		}, nil
	}

	return func(r *binary.StreamReader, ptr unsafe.Pointer) (err error) {
		head := (*sliceHeader)(ptr)
		calcSize := int(r.ReadUvarint())

		if calcSize > head.cap {
			return errors.New("not enough capacity in slice")
		}

		for i := 0; i < calcSize; i++ {
			elemPtr := unsafe.Add(head.data, uintptr(i)*itemSize)

			if err = dec(r, elemPtr); err != nil {
				return err
			}
		}

		head.len = calcSize

		return
	}, nil
}
