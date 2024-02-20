package bin

import (
	"reflect"
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

type arrayType struct {
	typ     Type
	typSize uintptr
	offset  uintptr
	len     int
}

func getArrayType(typ reflect.Type, offset uintptr, hasher func(reflect.Kind)) (Type, error) {
	l := typ.Len()
	hasher(reflect.Kind(l))

	elem := typ.Elem()
	subtyp, err := getType(elem, 0, hasher)

	if err != nil {
		return nil, err
	}

	t := arrayType{
		typ:     subtyp,
		typSize: elem.Size(),
		offset:  offset,
		len:     l,
	}

	return t, nil
}

func (t arrayType) EncodedSize(ptr unsafe.Pointer) (s int) {
	for i := 0; i < t.len; i++ {
		s += t.typ.EncodedSize(unsafe.Add(ptr, t.offset+uintptr(i)*t.typSize))
	}

	return
}

func (t arrayType) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	for i := 0; i < t.len; i++ {
		t.typ.encode(unsafe.Add(ptr, t.offset+uintptr(i)*t.typSize), b)
	}
}

func (t arrayType) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) (err error) {
	for i := 0; i < t.len; i++ {
		t.typ.decode(unsafe.Add(ptr, t.offset+uintptr(i)*t.typSize), b)
	}

	return
}
