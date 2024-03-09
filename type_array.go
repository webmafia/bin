package bin

import (
	"reflect"
	"unsafe"

	"github.com/webmafia/fast"
)

type arrayType struct {
	typ     Type
	typSize uintptr
	offset  uintptr
	len     int
}

func getArrayType(typ reflect.Type, offset uintptr) (Type, error) {
	elem := typ.Elem()
	subtyp, err := getType(elem, 0)

	if err != nil {
		return nil, err
	}

	t := arrayType{
		typ:     subtyp,
		typSize: elem.Size(),
		offset:  offset,
		len:     typ.Len(),
	}

	return t, nil
}

func (t arrayType) encodedSize(ptr unsafe.Pointer) (s int) {
	for i := 0; i < t.len; i++ {
		s += t.typ.encodedSize(unsafe.Add(ptr, t.offset+uintptr(i)*t.typSize))
	}

	return
}

func (t arrayType) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	for i := 0; i < t.len; i++ {
		t.typ.encode(unsafe.Add(ptr, t.offset+uintptr(i)*t.typSize), b)
	}
}

func (t arrayType) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader, nocopy bool) (err error) {
	for i := 0; i < t.len; i++ {
		t.typ.decode(unsafe.Add(ptr, t.offset+uintptr(i)*t.typSize), b, nocopy)
	}

	return
}
