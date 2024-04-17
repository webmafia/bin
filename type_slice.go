package bin

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/webmafia/fast/binary"
)

type sliceType struct {
	typ              Type
	typSize          uintptr
	offset           uintptr
	refTyp           reflect.Type
	allowAllocations bool
}

type sliceHeader struct {
	data unsafe.Pointer
	len  int
	cap  int
}

func getSliceType(typ reflect.Type, offset uintptr, opt *CoderOptions) (Type, error) {
	elem := typ.Elem()
	subtyp, err := getType(elem, 0, opt)

	if err != nil {
		return nil, err
	}

	t := sliceType{
		typ:              subtyp,
		typSize:          elem.Size(),
		offset:           offset,
		refTyp:           typ,
		allowAllocations: opt.AllowAllocations,
	}

	return t, nil
}

func (t sliceType) encodedSize(ptr unsafe.Pointer) (s int) {
	head := t.head(ptr)
	s += sizeUvarint(uint64(head.len))

	for i := 0; i < head.len; i++ {
		t.typ.encodedSize(unsafe.Add(head.data, uintptr(i)*t.typSize))
	}

	return
}

func (t sliceType) encode(ptr unsafe.Pointer, b binary.Writer) {
	head := t.head(ptr)
	b.WriteUvarint(uint64(head.len))

	for i := 0; i < head.len; i++ {
		t.typ.encode(unsafe.Add(head.data, uintptr(i)*t.typSize), b)
	}
}

func (t sliceType) decode(ptr unsafe.Pointer, b binary.Reader, nocopy bool) (err error) {
	head := t.head(ptr)
	calcSize := int(b.ReadUvarint())

	if calcSize > head.cap {
		if !t.allowAllocations {
			return errors.New("not enough capacity in slice")
		}

		// Allocate new slice with the calculated size
		newSlice := reflect.MakeSlice(t.refTyp, calcSize, calcSize)
		head.data = newSlice.UnsafePointer()
		head.cap = calcSize
		head.len = 0
	}

	for i := 0; i < calcSize; i++ {
		elemPtr := unsafe.Add(head.data, uintptr(i)*t.typSize)
		if err = t.typ.decode(elemPtr, b, nocopy); err != nil {
			return err
		}
	}

	head.len = calcSize

	return
}

//go:inline
func (t sliceType) head(ptr unsafe.Pointer) *sliceHeader {
	return (*sliceHeader)(unsafe.Add(ptr, t.offset))
}
