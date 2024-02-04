package bin

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

type sliceType struct {
	typ     Type
	typSize uintptr
	offset  uintptr
}

type sliceHeader struct {
	data unsafe.Pointer
	len  int
	cap  int
}

func getSliceType(typ reflect.Type, offset uintptr, hasher func(reflect.Kind)) (Type, error) {
	elem := typ.Elem()
	subtyp, err := getType(elem, 0, hasher)

	if err != nil {
		return nil, err
	}

	t := sliceType{
		typ:     subtyp,
		typSize: elem.Size(),
		offset:  offset,
	}

	return t, nil
}

func (t sliceType) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	head := (*sliceHeader)(unsafe.Add(ptr, t.offset))
	b.WriteUvarint(uint64(head.len))

	for i := 0; i < head.len; i++ {
		t.typ.encode(unsafe.Add(head.data, uintptr(i)*t.typSize), b)
	}
}

func (t sliceType) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) (err error) {
	head := (*sliceHeader)(unsafe.Add(ptr, t.offset))
	calcSize := head.len + int(b.ReadUvarint())

	if calcSize > head.cap {
		return errors.New("not enough capacity in slice")
	}

	for i := head.len; i < calcSize; i++ {
		t.typ.decode(unsafe.Add(head.data, uintptr(i)*t.typSize), b)
	}

	head.len = calcSize

	return
}
