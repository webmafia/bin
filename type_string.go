package bin

import (
	"unsafe"

	"github.com/webmafia/fast/binary"
)

type stringType struct {
	offset uintptr
}

func (t stringType) encodedSize(ptr unsafe.Pointer) (s int) {
	str := *(*string)(unsafe.Add(ptr, t.offset))
	l := len(str)
	s += sizeUvarint(uint64(l))
	s += l

	return
}

func (f stringType) encode(ptr unsafe.Pointer, b binary.Writer) {
	str := *(*string)(unsafe.Add(ptr, f.offset))
	b.WriteUvarint(uint64(len(str)))
	b.WriteString(str)
}

func (f stringType) decode(ptr unsafe.Pointer, b binary.Reader, nocopy bool) error {
	n := b.ReadUvarint()

	if nocopy {
		*(*string)(unsafe.Add(ptr, f.offset)) = b.ReadString(int(n))
	} else {
		*(*string)(unsafe.Add(ptr, f.offset)) = string(b.ReadBytes(int(n)))
	}

	return nil
}
