package bin

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/webmafia/fast"
)

type sliceType struct {
	typ              Type
	typSize          uintptr
	offset           uintptr
	allowAllocations bool
}

type sliceHeader struct {
	data unsafe.Pointer
	len  int
	cap  int
}

func getSliceType(typ reflect.Type, offset uintptr, allowAllocations bool) (Type, error) {
	elem := typ.Elem()
	subtyp, err := getType(elem, 0, allowAllocations)

	if err != nil {
		return nil, err
	}

	t := sliceType{
		typ:              subtyp,
		typSize:          elem.Size(),
		offset:           offset,
		allowAllocations: allowAllocations,
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

func (t sliceType) encode(ptr unsafe.Pointer, b fast.Writer) {
	head := t.head(ptr)
	b.WriteUvarint(uint64(head.len))

	for i := 0; i < head.len; i++ {
		t.typ.encode(unsafe.Add(head.data, uintptr(i)*t.typSize), b)
	}
}

func (t sliceType) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader, nocopy bool) (err error) {
	head := t.head(ptr)
	calcSize := head.len + int(b.ReadUvarint())

	if calcSize > head.cap {
		if !t.allowAllocations {
			return errors.New("not enough capacity in slice")
		}

		oldBytes := *(*[]byte)(unsafe.Pointer(&sliceHeader{
			data: head.data,
			len:  head.len * int(t.typSize),
			cap:  head.len * int(t.typSize),
		}))

		newBytes := fast.MakeNoZero(calcSize * int(t.typSize))
		copy(newBytes, oldBytes)

		newBytesHead := (*sliceHeader)(unsafe.Pointer(&newBytes))
		head.data = newBytesHead.data
		head.cap = calcSize
	}

	for i := head.len; i < calcSize; i++ {
		t.typ.decode(unsafe.Add(head.data, uintptr(i)*t.typSize), b, nocopy)
	}

	head.len = calcSize

	return
}

//go:inline
func (t sliceType) head(ptr unsafe.Pointer) *sliceHeader {
	return (*sliceHeader)(unsafe.Add(ptr, t.offset))
}
